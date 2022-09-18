package storage

import "github.com/dollarkillerx/go-proxy-manager/internal/app/backend/pkg/models"

type Interface interface {
	GetUser(account string) (*models.UserCenter, error)
	UpdateUser(account string, password string) error

	GetAllTask() ([]models.Tasks, error)
	RegisterTask(task models.Tasks) error
	UpdateTasks(task models.Tasks) error
	DeleteTasks(taskID string) error
}
