package databases

import (
	"cleancode/config"
	"cleancode/models"
)

func GetAllUsers() (interface{}, error) {
	users := []models.User{}
	err := config.Db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetSingleUser(userId int) (interface{}, error) {
	user := models.User{}
	err := config.Db.Find(&user, userId).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}
