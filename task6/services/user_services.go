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
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)


var (
	adminCreated bool
	
)


func hashPassword(password string) (string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil{
		return "", errors.New("couldnt Generate Password")
	}
	return string(bytes), nil
}

func verifyPassword(existingPassword string, newPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(existingPassword), []byte(newPassword))

	return err == nil
}



func CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Second)
	filter := bson.D{{Key: "Email", Value: user.Email}}

	userCollection := GetDatabaseCollection("user")


	exisitingUser := userCollection.FindOne(ctx, filter)
	defer cancel()

	if exisitingUser != nil{
		return errors.New("user Already Exists")
	}

	if !adminCreated{
		adminCreated = true
		user.Role = Admin
	}

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAT, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.ID = primitive.NewObjectID()
	user.UserID = user.ID.Hex()
	password, err := hashPassword(*user.Password)
	if err != nil{
		return errors.New("couldn't Parse Password")
	}
	user.Password = &password

	token, refreshToken, err := Utils.GenerateTokens(*user.Email, *user.Name, user.UserID, *user.Role)
	if err != nil{
		return errors.New("internal server error")
	}

	user.Token = token 
	user.RefreshToken = refreshToken

	_, err = userCollection.InsertOne(ctx, user)
	if err != nil{
		return errors.New("internal server error")
	}
	defer cancel()
	return nil

	
}



func Login( user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 100)
	

	userCollection := GetDatabaseCollection("user")
	filter := bson.D{{Key:"email", Value:user.Email}}
	var existingUser models.User 
	err := userCollection.FindOne(ctx, filter).Decode(&existingUser)
	cancel() 

	if err != nil{
		return errors.New("user doesnot Exist")
	}

	if existingUser.Email == nil{
		return errors.New("user doesnot Exist")
	}

	ok := verifyPassword(*user.Password, *existingUser.Password)
	if !ok{
		return errors.New("wrong Password")
	}

	
	token, refreshToken, err := Utils.GenerateTokens(*user.Email, *user.Name, user.UserID, *user.Role)
	if err != nil{
		return errors.New("internal server error")
	}


	updateObj := Utils.UpdateAllTokens(*token, *refreshToken)
	upsert := true

	filters := bson.M{"user_id":existingUser.UserID}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err = userCollection.UpdateOne(
		ctx,
		filters,
		bson.D{
			{Key: "$set", Value: updateObj},
		},
		&opt,
	)

	defer cancel()

	if err!=nil{
		log.Panic(err)
		return errors.New("internal Server Error User not updated")
	}
	var newuser models.User
	err = userCollection.FindOne(ctx, bson.M{"UserID":existingUser.UserID}).Decode(&newuser)

		if err != nil {
				return errors.New("internal Server Error")
	}
	return nil

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
