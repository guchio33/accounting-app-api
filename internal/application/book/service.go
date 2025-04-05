package application

import (
	domain "accounting-app-api/internal/domain/book"
	infrastructure "accounting-app-api/internal/infrastructure/mysql"
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