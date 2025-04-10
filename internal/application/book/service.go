package application

import (
	domain "accounting-app-api/internal/domain/book"
	infrastructure "accounting-app-api/internal/infrastructure/mysql"
	"fmt"
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

func (s *BookService) GetBook(bookId int) (*domain.Book, error) {
	return s.Repo.GetBook(bookId)
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

func (s *BookService) UpdateBook(id int, title *string, authour *string) (error) {
	// idから情報を取得
	book, err := s.Repo.GetBook(id)
	if err != nil {
		return fmt.Errorf("failed to get book: %w", err)
	}

	if title != nil {
		book.Title = *title 
	}
	if authour != nil {
		book.Author = *authour
	}

	return s.Repo.UpdateBook(book, id)
}