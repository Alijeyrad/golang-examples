package services

import (
	"library/models"
	"library/repositories"
)

type BookService interface {
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id uint) (*models.Book, error)
	CreateBook(book *models.Book) error
	UpdateBook(book *models.Book) error
	DeleteBook(id uint) error
}

type bookService struct {
	repo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{repo: repo}
}

func (s *bookService) GetAllBooks() ([]models.Book, error) {
	return s.repo.GetAll()
}

func (s *bookService) GetBookByID(id uint) (*models.Book, error) {
	return s.repo.GetByID(id)
}

func (s *bookService) CreateBook(book *models.Book) error {
	return s.repo.Create(book)
}

func (s *bookService) UpdateBook(book *models.Book) error {
	return s.repo.Update(book)
}

func (s *bookService) DeleteBook(id uint) error {
	book, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(book)
}
