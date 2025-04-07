package domain

type BookRepository interface {
	// すべての本を取得する
	GetAllBooks() ([]Book, error)

	// 本を追加する
	AddBook(book *Book) error

	//本を削除
	DeleteBook(bookId int) error
}