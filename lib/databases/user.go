package databases

import (
	"cleancode/config"
	"cleancode/models"
)

func GetAllUsers() (interface{}, int, error) {
	users := []models.User{}
	result := config.Db.Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if result.RowsAffected > 0 {
		return users, 1, nil
	}

	return "user data not found", 0, nil
}

func GetSingleUser(userId int) (interface{}, int, error) {
	user := models.User{}
	result := config.Db.Find(&user, userId)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if result.RowsAffected > 0 {
		return user, 1, nil
	}

	return "user data not found", 0, nil
}

func CreateNewUser(user *models.User) (interface{}, error) {
	result := config.Db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func DeleteUser(userId int) (interface{}, int, error) {
	user := models.User{}
	result := config.Db.Delete(&user, userId)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if result.RowsAffected > 0 {
		return "deleted", 1, nil
	}
	return "user data not found", 0, nil
}

func UpdateUser(userId int, newUser models.User) (interface{}, int, error) {
	user := models.User{}
	findResult := config.Db.Find(&user, userId)
	if findResult.Error != nil {
		return nil, 0, findResult.Error
	}

	if findResult.RowsAffected > 0 {
		updatedResult := config.Db.Model(&user).Updates(newUser)
		if updatedResult.Error != nil {
			return nil, 0, updatedResult.Error
		}
		return user, 1, nil
	}

	return "user data not found", 0, nil
}
