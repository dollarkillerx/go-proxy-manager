package simple

import (
	"github.com/dollarkillerx/go-proxy-manager/internal/app/backend/pkg/models"
	"github.com/dollarkillerx/go-proxy-manager/internal/utils"

	"log"
)

func (s *Simple) GetUser(account string) (*models.UserCenter, error) {
	var us models.UserCenter
	err := s.db.Model(&models.UserCenter{}).Where("account = ?", account).First(&us).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &us, nil
}

func (s *Simple) UpdateUser(account string, password string) error {
	var us models.UserCenter
	err := s.db.Model(&models.UserCenter{}).First(&us).Error
	if err != nil {
		log.Println(err)
		return err
	}

	err = s.db.Model(&models.UserCenter{}).Where("1=1").Updates(map[string]interface{}{
		"account":  account,
		"password": utils.GenPassword(password, us.Sale),
	}).Error
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (s *Simple) GetAllTask() ([]models.Tasks, error) {
	var tasks []models.Tasks
	err := s.db.Model(&models.Tasks{}).Find(&tasks).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return tasks, nil
}

func (s *Simple) RegisterTask(task models.Tasks) error {
	err := s.db.Model(&models.Tasks{}).Create(&task).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *Simple) UpdateTasks(task models.Tasks) error {
	err := s.db.Model(&models.Tasks{}).
		Where("id = ?", task.ID).
		Updates(map[string]interface{}{
			"domain":       task.Domain,
			"enable_proxy": task.EnableProxy,
			"enable_ssl":   task.EnableSSL,
			"enable_waf":   task.EnableWaf,
			"proxy_type":   task.ProxyType,
			"payload":      task.Payload,
			"certificate":  task.CreatedAt,
		}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *Simple) DeleteTasks(taskID string) error {
	err := s.db.Model(&models.Tasks{}).Where("id = ?", taskID).Delete(&models.Tasks{}).Error
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
