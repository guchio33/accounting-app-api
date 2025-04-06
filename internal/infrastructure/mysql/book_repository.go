package infrastructure

import (
	domain "accounting-app-api/internal/domain/book"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

// BookRepositoryはBookデータをMySQLから操作するリポジトリ
type BookRepository struct {
	DB *sqlx.DB
}

// 新しいBookRepositoryのインスタンスを作成します
func NewBookRepository(db *sqlx.DB) *BookRepository {
	return &BookRepository{DB: db}
}

func (r *BookRepository) GetAllBooks() ([]domain.Book, error) {
	var books []domain.Book
	query := "SELECT * from books"
	err := r.DB.Select(&books, query)
	log.Print("bookの返却値: %+v",books)
	log.Print("fagggaga")

	if err != nil {
		return nil, fmt.Errorf("failed to get books: %w", err)
	}

	return books, nil
}