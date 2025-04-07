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
	log.Print(query, &books)

	if err != nil {
		return nil, fmt.Errorf("failed to get books: %w", err)
	}

	return books, nil
}

func (r *BookRepository) GetBook(bookId int) (*domain.Book, error) {
	var book domain.Book

	query := "SELECT * from books WHERE id = ? "
	err := r.DB.Get(&book, query, bookId)

	if err != nil {
		return nil, fmt.Errorf("failed to get books: %w", err)
	}

	return &book, nil
}

func (r *BookRepository) AddBook(book *domain.Book) error {
	query := "INSERT INTO books (title, author, created_at, updated_at) VALUES(?,?,?,?)"
	_,err := r.DB.Exec(query, book.Title, book.Author, book.CreatedAt, book.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to add book: %w", err)
	}

	return nil
}

func (r *BookRepository) DeleteBook(bookId int) error {
	query := "DELETE FROM books WHERE id =?"
	_,err := r.DB.Exec(query, bookId)

	if err != nil {
		return fmt.Errorf("failed to delete book: %w", err)
	}

	return nil
}

func (r *BookRepository) UpdateBook(book *domain.Book, bookId int) error {
	query := "Update books SET title=?, author=? WHERE id =?"
	_,err := r.DB.Exec(query, book.Title, book.Author, bookId)

	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}

	return nil
}