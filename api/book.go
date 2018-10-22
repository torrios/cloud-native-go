package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
	Description string `json:"description,omitempty"`
}

func (book *Book) ToJSON() []byte {
	jsonBytes, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}

	return jsonBytes
}

func FromJSON(data []byte) Book {
	newBook := Book{}
	err := json.Unmarshal(data, &newBook)
	if err != nil {
		panic(err)
	}

	return newBook
}

var books = map[string]Book{
	"12234567" : {Title: "Fear and Loathing in Las Vegas", Author: "Hunter Thomson", ISBN: "12234567"},
	"23375746" : {Title: "The Art of Heckie", Author: "Hector Rios", ISBN: "23375746", Description: "A book dedicated to the amazingness of Heckie"},
}

func BooksHandleFuncMarshall(w http.ResponseWriter, r *http.Request) {
	booksData, err := json.Marshal(books)

	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf8")
	w.Write(booksData)
}

func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		allBooks := AllBooks()
		writeJSON(w, allBooks)
	case http.MethodPost:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		newBook := FromJSON(body)
		isbn, created := CreateBook(newBook)
		if created {
			w.Header().Add("Location", "/api/books/" + isbn)
			w.WriteHeader(http.StatusCreated)
		} else {
			w.WriteHeader(http.StatusConflict)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method."))
	}
}

func BookHandleFunc(w http.ResponseWriter, r *http.Request) {
	isbn := r.URL.Path[len("/api/books/"):]
	switch method := r.Method; method {
	case http.MethodGet:
		book, found := GetBook(isbn)
		if found {
			writeJSON(w, book)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPut:
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		updatedBook := FromJSON(body)
		exists := UpdateBook(isbn, updatedBook)
		if exists {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodDelete:
		DeleteBook(isbn)
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method"))
	}
}

func writeJSON(w http.ResponseWriter, data interface{}) {
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.Header().Add("Content-Type", "application/json; charset=utf8")
	w.Write(b)
}

//AllBooks returns a slice of all the books
func AllBooks() []Book {
	values := make([]Book, len(books))
	idx := 0
	for _, bookVal := range books {
		values[idx] = bookVal
		idx++
	}

	return values
}

func CreateBook(newBook Book) (string, bool) {
	_, exists := books[newBook.ISBN]
	if exists {
		return "", false
	}

	books[newBook.ISBN] = newBook
	return newBook.ISBN, true
}

func GetBook(isbn string) (Book, bool) {
	bookVal, exists := books[isbn]
	return bookVal, exists
}

func UpdateBook(isbn string, book Book) bool {
	_, exists := books[isbn]
	if exists {
		books[isbn] = book
	}
	return exists
}

func DeleteBook(isbn string) {
	delete(books, isbn)
}
