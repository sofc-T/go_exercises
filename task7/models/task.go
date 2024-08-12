package models

import(
 	"context"
)




type Task struct{
	Id int  `json:"_id"`
	Title string  `json:"title"`
	TaskId int  `json:"task_id"`
}

type TaskRepository interface {
	FetchTasks(ctx context.Context) ([]Task, error )
	FindTask(ctx context.Context, id int) (Task , error)
	UpdateTask(ctx context.Context, id int, title string) (Task, error)
	DeleteTask(ctx context.Context, id int) (error)
	InsertTask(ctx context.Context,task Task) (Task , error)

}

type TaskUsecase interface {
	Fetch(ctx context.Context) ([]Task, error )
	Find(ctx context.Context, id int) (Task , error)
	Update(ctx context.Context, id int, title string) (Task, error)
	Delete(ctx context.Context, id int) (error)
	Create(ctx context.Context,task Task) (Task , error)
}

