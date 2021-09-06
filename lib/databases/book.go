package databases

import (
	"cleancode/config"
	"cleancode/models"
)

func GetAllBooks() (interface{}, int, error) {
	books := []models.Book{}
	result := config.Db.Find(&books)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if result.RowsAffected > 0 {
		return books, 1, nil
	}

	return "Book not found", 0, nil
}

func GetSingleBook(bookId int) (interface{}, int, error) {
	book := models.Book{}
	result := config.Db.Find(&book, bookId)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if result.RowsAffected > 0 {
		return book, 1, nil
	}

	return "Book not found", 0, nil
}

func CreateNewBook(book *models.Book) (interface{}, error) {
	result := config.Db.Create(&book)
	if result.Error != nil {
		return nil, result.Error
	}

	return book, nil
}

func DeleteBook(bookId int) (interface{}, int, error) {
	book := models.Book{}
	result := config.Db.Delete(&book, bookId)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	if result.RowsAffected > 0 {
		return "deleted", 1, nil
	}
	return "Book not found", 0, nil
}

func UpdateBook(bookId int, newBook models.Book) (interface{}, int, error) {
	book := models.Book{}
	findResult := config.Db.Find(&book, bookId)
	if findResult.Error != nil {
		return nil, 0, findResult.Error
	}

	if findResult.RowsAffected > 0 {
		updatedResult := config.Db.Model(&book).Updates(newBook)
		if updatedResult.Error != nil {
			return nil, 0, updatedResult.Error
		}
		return book, 1, nil
	}

	return "Book not found", 0, nil
}
