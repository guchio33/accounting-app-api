package handler

import (
	application "accounting-app-api/internal/application/book"
	"encoding/json"
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