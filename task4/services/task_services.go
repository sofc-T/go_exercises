package services 

import (
	"github.com/sofc-t/task_manager/models"
)

type TaskManager struct{
	Tasks []*models.Task
}

type TaskManagerInterface interface{
	GetAllTasks() []models.Task 
	GetTask(id int) models.Task
	UpdateTask(id int, task models.Task) models.Task 
	DeleteTask(id int)
	CreateTask(id int, Title string)
}


func (t *TaskManager) GetAllTasks() []*models.Task{
	return t.Tasks
}

func (t *TaskManager)  GetTask(id int) *models.Task{
	for _, task := range t.Tasks{
		if task.Id == id{
			return task
		}
	}
	return nil
}

func (t *TaskManager) UpdateTask(id int, title string, ) *models.Task{
	for _, task := range(t.Tasks){
		if task.Id == id{
			task.Title = title
			return task
		}
	}
	return nil
}

func (t *TaskManager) DeleteTask(id int) map[string]string {
	for idx, task := range(t.Tasks){
		if task.Id == id{
			t.Tasks = append(t.Tasks[:idx], t.Tasks[idx + 1:]...)		
			message := map[string]string {"message" : "Succesfully Deleted"}
			return message
		}
	}
	return nil
}


func (t *TaskManager) CreateTask(id int, title string) *models.Task{
	task := &models.Task{Id: id, Title: title}
	t.Tasks = append(t.Tasks, task)
	return task
}