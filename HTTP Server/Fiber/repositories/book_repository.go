package repositories

import (
	"yourapp/config"
	"yourapp/models"
)

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetByID(id uint) (*models.Book, error)
	Create(book *models.Book) error
	Update(book *models.Book) error
	Delete(book *models.Book) error
}

type bookRepository struct{}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

func (r *bookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := config.GetDB().Find(&books).Error
	return books, err
}

func (r *bookRepository) GetByID(id uint) (*models.Book, error) {
	var book models.Book
	err := config.GetDB().First(&book, id).Error
	return &book, err
}

func (r *bookRepository) Create(book *models.Book) error {
	return config.GetDB().Create(book).Error
}

func (r *bookRepository) Update(book *models.Book) error {
	return config.GetDB().Save(book).Error
}

func (r *bookRepository) Delete(book *models.Book) error {
	return config.GetDB().Delete(book).Error
}
