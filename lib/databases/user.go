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

func CreateNewUser(user *models.User) (interface{}, error) {
	err := config.Db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(userId int) (interface{}, error) {
	user := models.User{}
	err := config.Db.Delete(&user, userId).Error
	if err != nil {
		return nil, err
	}

	return "deleted", nil
}

func UpdateUser(userId int, newUser models.User) (interface{}, error) {
	user := models.User{}
	notFoundId := config.Db.Find(&user, userId).Error
	if notFoundId != nil {
		return nil, notFoundId
	}

	err := config.Db.Model(&user).Updates(newUser).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
