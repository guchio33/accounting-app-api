package handler

import (
	application "accounting-app-api/internal/application/book"
	"encoding/json"
	"fmt"
	"net/http"
)

type BookHandler struct {
	Service *application.BookService
}

func NewBookHandler(service *application.BookService) *BookHandler {
	return &BookHandler{Service: service}
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.Service.GetAllBooks()
	if err != nil {
		http.Error(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	// レスポンスヘッダーにセット
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// HTTPリクエストを受け付けて書籍を追加する処理
func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request){
	//　書籍を定義
	var book struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	// HTTPリクエストのボディからJSONデータを読み取り、デコード
	// 失敗した場合return
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	fmt.Println(book)

	if err := h.Service.AddBook(book.Title, book.Author); err != nil {
		http.Error(w, "Failed to add book", http.StatusInternalServerError)
		return
	}
	
	//ステータスコードを記述
	w.WriteHeader(http.StatusCreated)
}