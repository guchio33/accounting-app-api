package handler

import (
	application "accounting-app-api/internal/application/book"
	"encoding/json"
	"net/http"
	"strconv"
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

	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}

	// レスポンスヘッダーにセット
	w.Header().Set("Content-Type", "application/json")
}


func (h *BookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	// URLからid取得
	id := r.PathValue("id")
	// idが存在しなかった場合、数字ではなかった場合
	intBookId, parseErr := strconv.Atoi(id)
	if parseErr != nil {
		http.Error(w, "ID must be a number or ID is missing ", http.StatusBadRequest)
    return
	}

	book, err := h.Service.GetBook(intBookId); 
	if err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	// ステータスコード200 OKを設定
	w.WriteHeader(http.StatusOK)
	// 書籍情報をJSONとしてレスポンスに返す
	json.NewEncoder(w).Encode(book)
}


// HTTPリクエストを受け付けて書籍を追加する処理
func (h *BookHandler) AddBook(w http.ResponseWriter, r *http.Request){
	//　書籍を定義
	var book struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	// HTTPリクエストのボディからJSONデータを読み取り、デコード
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.AddBook(book.Title, book.Author); err != nil {
		http.Error(w, "Failed to add book", http.StatusInternalServerError)
		return
	}
	
	//ステータスコードを記述
	w.WriteHeader(http.StatusCreated)
}


func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// URLからid取得
	id := r.PathValue("id")

	// idが存在しなかった場合、数字ではなかった場合
	intBookId, parseErr := strconv.Atoi(id)
	if parseErr != nil {
		http.Error(w, "ID must be a number or ID is missing ", http.StatusBadRequest)
    return
	}

	if err := h.Service.DeleteBook(intBookId); err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	//ステータスコードを記述
	w.WriteHeader(http.StatusOK)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	// URLからid取得
	id := r.PathValue("id")

	// *string型とすることで、nullを判別することができる
	var book struct {
		Title  *string `json:"title"`
		Author *string `json:"author"`
	}

	// idが存在しなかった場合、数字ではなかった場合
	intBookId, parseErr := strconv.Atoi(id)
	if parseErr != nil {
		http.Error(w, "ID must be a number or ID is missing ", http.StatusBadRequest)
    return
	}

	// HTTPリクエストのボディからJSONデータを読み取り、デコード
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.UpdateBook(intBookId, book.Title, book.Author); err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	//ステータスコードを記述
	w.WriteHeader(http.StatusOK)
}