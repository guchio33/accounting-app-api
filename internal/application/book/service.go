package application

import (
	domain "accounting-app-api/internal/domain/book"
	infrastructure "accounting-app-api/internal/infrastructure/mysql"
	"time"
)

// Bookリポジトリを使ってビジネスロジックを提供します
type BookService struct {
	Repo *infrastructure.BookRepository
}

// BookServiceのインスタンスを作成します
func NewBookService(repo *infrastructure.BookRepository) *BookService {
	return &BookService{Repo: repo}
}

func (s *BookService) GetAllBooks() ([]domain.Book, error) {
	return s.Repo.GetAllBooks()
}

func (s *BookService) AddBook(title string, author string) error {
	book := &domain.Book{
		Title:     title,
		Author:    author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return s.Repo.AddBook(book)
}

func (s *BookService) DeleteBook(bookId int) error {
	return s.Repo.DeleteBook(bookId)
}