# Books-API-golang

This is a simple RESTful API for managing books built using Go's `net/http` package. It allows users to add, update, delete, and list books in a JSON format.

## Features

- Add a new book (POST /books)
- List all books (GET /books)
- Update an existing book (PUT /books/{id})
- Delete a book (DELETE /books/{id})

## Technologies Used

- Go (Golang)
- net/http package for handling HTTP requests and responses
- JSON for data interchange

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/AymanJanati/books-api-golang
   ```
2. Change directory to the project folder:
   ```bash
   cd books-api-golang
   ```
3. Run the server:
   ```bash
   go run main.go
   ```

The server will start at `http://localhost:8080`.

## Usage

### Adding a Book
To add a book, send a POST request to `/books` with a JSON body:
```json
{
  "title": "Book Title",
  "author": "Book Author"
}
```

### Listing Books
To list all books, send a GET request to `/books`.

### Updating a Book
To update a book, send a PUT request to `/books/{id}` with a JSON body:
```json
{
  "title": "Updated Title",
  "author": "Updated Author"
}
```

### Deleting a Book
To delete a book, send a DELETE request to `/books/{id}`.
