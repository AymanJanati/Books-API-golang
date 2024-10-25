package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books = []Book{}
var idCounter int = 1

func addBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	var newBook Book

	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		http.Error(w, "invalide input", http.StatusBadRequest)
		return
	}

	newBook.Id = idCounter
	idCounter++
	Books = append(Books, newBook)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newBook)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Path[len("/books/"):])
	if err != nil {
		http.Error(w, "invalide book id!", http.StatusBadRequest)
		return
	}

	for i, book := range Books {
		if book.Id == id {
			Books = append(Books[:i], Books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "book not found", http.StatusNotFound)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Path[len("/books/"):])
	if err != nil {
		http.Error(w, "Invalid book ID!", http.StatusBadRequest)
		return
	}

	var updatedBook Book
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	for i, book := range Books {
		if book.Id == id {
			updatedBook.Id = id
			Books[i] = updatedBook
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}

	http.Error(w, "Book not found", http.StatusNotFound)
}

func listBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if len(Books) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	limitstr := r.URL.Query().Get("limit")
	offsetstr := r.URL.Query().Get("offset")

	if limitstr == "" || offsetstr == "" {
		http.Error(w, "Invalide input: limit and offset missing", http.StatusBadRequest)
		return
	}

	l, err := strconv.Atoi(limitstr)
	if err != nil || l <= 0 {
		http.Error(w, "Invalide input: limite must be positive.", http.StatusBadRequest)
		return
	}
	of, err := strconv.Atoi(offsetstr)
	if err != nil || of < 0 {
		http.Error(w, "Invalide input: offset must be non-negative", http.StatusBadRequest)
		return
	}

	end := of + l
	if end > len(Books) {
		end = len(Books)
	}

	retBooks := Books[of:end]

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(retBooks)
}

func main() {
	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			listBooks(w, r)
		} else {
			addBook(w, r)
		}
	})
	http.HandleFunc("/books/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			updateBook(w, r)
		} else {
			deleteBook(w, r)
		}
	})

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
