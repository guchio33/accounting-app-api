package domain

import "time"

// bookの構造対
type Book struct {
	ID 						int64 		`db:"id" json:"id"`
	Title         string    `db:"title" json:"title"`
	Author        string    `db:"author" json:"author"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}