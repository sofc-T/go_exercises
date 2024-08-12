package usecases

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/sofc-t/task_manager/task7/models"
	Utils "github.com/sofc-t/task_manager/task7/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepositry models.UserRepository
	timeOut time.Duration
}
var (
	adminCreated bool
	adm  = "admin"
	gue = "user"
	admin = &adm 
	guest = &gue
)



func NewUserUsecase(userRepositry models.UserRepository, timeOut time.Duration) *userUsecase {
	return &userUsecase{
		userRepositry: userRepositry,
		timeOut: timeOut,
	}
}

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

func(u userUsecase) Create(ctx context.Context, user models.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()

	if !adminCreated {
		adminCreated = true
		user.Role = admin
		log.Println("Assigned Admin role to new user")
	} else {
		user.Role = guest
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

	return u.userRepositry.CreateUser(ctx, user)

}

func(u userUsecase) Login(ctx context.Context, user models.User) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()
	return u.userRepositry.Login(ctx, user)

}

func(u userUsecase) FetchAll(ctx context.Context) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()
	return u.userRepositry.FetchAllUsers(ctx)

}


func(u userUsecase) FetchById(ctx context.Context, id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()
	return u.userRepositry.FetchByID(ctx, id)

}



func(u userUsecase) PromoteUser(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()
	return u.userRepositry.PromoteUser(ctx, id)

}






