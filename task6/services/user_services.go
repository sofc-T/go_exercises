package services

import (
	"context"
	"errors"
	"log"
	"time"

	
	"github.com/sofc-t/task_manager/task6/models"
	"github.com/sofc-t/task_manager/task6/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)


var (
	adminCreated bool
)


func hashPassword(password string) (string, error){
	log.Println("hashing", password)
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil{
		log.Println("couldnt Generate Password")
		return "", errors.New("couldnt Generate Password")
	}
	log.Println("hashed")
	return string(bytes), nil
}

func verifyPassword(existingPassword string, newPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(newPassword), []byte(existingPassword))
	if err != nil {
        log.Println("Password verification failed:", err)
    } else {
        log.Println("Password verification successful")
    }

	return err == nil
}


func CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.D{{Key: "Email", Value: user.Email}}
	userCollection := GetDatabaseCollection("user")

	var existingUser models.User
	err := userCollection.FindOne(ctx, filter).Decode(&existingUser)
	if err == nil {
		log.Println("User already exists with email:", user.Email)
		return errors.New("user Already Exists")
	} else if err != mongo.ErrNoDocuments {
		log.Println("Error checking for existing user:", err)
		return errors.New("internal server error")
	}
	log.Println("No existing user found")

	if !adminCreated {
		adminCreated = true
		user.Role = Admin
		log.Println("Assigned Admin role to new user")
	} else {
		user.Role = Guest
		log.Println("Assigned role : user")
	}

	user.CreatedAt = time.Now()
	user.UpdatedAT = time.Now()
	user.ID = primitive.NewObjectID()
	user.UserID = user.ID.Hex()
	log.Println("id generated")

	password, err := hashPassword(*user.Password)
	if err != nil {
		log.Println("Error hashing password:", err)
		return errors.New("couldn't parse password")
	}
	user.Password = &password

	token, refreshToken, err := Utils.GenerateTokens(*user.Email, *user.Name, user.UserID, *user.Role)
	if err != nil {
		log.Println("Error generating tokens:", err)
		return errors.New("internal server error")
	}

	user.Token = token
	user.RefreshToken = refreshToken

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		log.Println("Error inserting user into database:", err)
		return errors.New("internal server error")
	}

	log.Println("User created successfully with email:", user.Email)
	return nil
}



func Login(user *models.User) (string, error ) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 100)
	defer cancel()

	userCollection := GetDatabaseCollection("user")
	filter := bson.D{{Key: "email", Value: user.Email}}
	var existingUser models.User

	log.Println("Attempting to find user with email:", user.Email)
	err := userCollection.FindOne(ctx, filter).Decode(&existingUser)
	if err != nil {
		log.Println("Error finding user:", err)
		return "", errors.New("user does not exist")
	}

	if existingUser.Email == nil {
		log.Println("User email is nil, user does not exist")
		return "", errors.New("user does not exist")
	}
	if existingUser.Email == nil || existingUser.Name == nil || existingUser.Role == nil {
		
		log.Println("One or more required fields are nil", user.Name, user.Email, user.Role, user)
		return "", errors.New("invalid user data")
	}

	log.Println("Verifying password for user:", user.Email)
	ok := verifyPassword(*user.Password, *existingUser.Password)
	if !ok {
		log.Println("Password verification failed for user:", user.Email)
		return "", errors.New("wrong password")
	}

	log.Println("Generating tokens for user:", user.Email)
	token, refreshToken, err := Utils.GenerateTokens(*existingUser.Email, *existingUser.Name, existingUser.UserID, *existingUser.Role)
	if err != nil {
		log.Println("Error generating tokens:", err)
		return "", errors.New("internal server error")
	}

	updateObj := Utils.UpdateAllTokens(*token, *refreshToken)
	upsert := true
	filters := bson.M{"user_id": existingUser.UserID}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	log.Println("Updating tokens for user with UserID:", existingUser.UserID)
	_, err = userCollection.UpdateOne(
		ctx,
		filters,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
		&opt,
	)

	if err != nil {
		log.Panic("Error updating user tokens:", err)
		return "", errors.New("internal server error, user not updated")
	}

	log.Println("Fetching updated user data for UserID:", existingUser.UserID)
	// var newuser models.User
	// err = userCollection.FindOne(ctx, bson.M{"UserID": existingUser.UserID}).Decode(&newuser)
	// if err != nil {
	// 	log.Println("Error fetching updated user data:", err)
	// 	return errors.New("internal server error")
	// }

	// log.Println("Login successful for user:", user.Email)
	return *token, nil
}




func FetchAllUsers() error {
	userCollection := GetDatabaseCollection("user")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := userCollection.Find(ctx, bson.D{})
    if err != nil {
        return errors.New("couldnt Fetch Data")
    }
    defer cursor.Close(ctx)

    var users []models.User
    if err = cursor.All(ctx, &users); err != nil {
		return errors.New("couldnt Parse Data")
    }

    return nil
}

func FetchUserByID(id string) (models.User, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userCollection := GetDatabaseCollection("user")
	var user models.User 

	iD, err := primitive.ObjectIDFromHex(id)
	if err != nil{
		return user, errors.New("invalid ID")
	}
	
    err = userCollection.FindOne(ctx, bson.M{"_id": iD}).Decode(&user)
	defer cancel()
	
    if err != nil {
        return user, errors.New("error User not found")
    }

	return user, nil
}


func PromoteUser(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userCollection := GetDatabaseCollection("user")
	
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}
	var user models.User
	
	err = userCollection.FindOne(ctx, bson.M{"userid": userID}).Decode(&user)
	if err != nil {
		return errors.New("user not found")
	}

	if user.Role == Admin {
		return errors.New("user is already an admin")
	}

	user.Role = Admin

	_, err = userCollection.UpdateOne(
		ctx,
		bson.M{"userid": userID},
		bson.D{
			{Key: "$set", Value: bson.M{"role": user.Role}},
		},
	)
	if err != nil {
		return errors.New("failed to promote user")
	}

	return nil
}
