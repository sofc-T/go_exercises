package usecases

import (
	"context"
	"time"

	"github.com/sofc-t/task_manager/task7/models"
)

type userUsecase struct {
	userRepositry models.UserRepository
	timeOut time.Duration
}



func NewUserUsecase(userRepositry models.UserRepository, timeOut time.Duration) *userUsecase {
	return &userUsecase{
		userRepositry: userRepositry,
		timeOut: timeOut,
	}
}


func(u userUsecase) Create(ctx context.Context, user models.User) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()
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






