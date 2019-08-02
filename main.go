package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Books struct {
	ID    string `json:"id"`
	ISBN  string `json:"isbn"`
	TITLE string `json:"title"`
}

var books []Books

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Books{})
}

func createBook(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open("mysql", "muntakim:pranto172472@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Connected Mysql server")
	w.Header().Set("Content-Type", "application/json")
	var book Books
	_ = json.NewDecoder(r.Body).Decode(&book)
	insert, err := db.Query(`INSERT INTO books (ISBN , TITLE ) VALUES(` + `"` + book.ISBN + `" ,` + `"` + book.TITLE + `")`)
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("User added to the database")
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "muntakim:pranto172472@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Connected Mysql server")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through books and find with id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			delete, err := db.Query("DELETE from books where ID=" + params["id"] + "")
			if err != nil {
				panic(err.Error())
			}
			defer delete.Close()
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "muntakim:pranto172472@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	//Loop through books and find with id
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Books
			_ = json.NewDecoder(r.Body).Decode(&book)
			update, err := db.Query(`UPDATE books set ISBN="` + book.ISBN + `",TITLE="` + book.TITLE + `" where ID=` + params["id"])
			if err != nil {
				panic(err.Error())
			}
			defer update.Close()
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	db, err := sql.Open("mysql", "muntakim:pranto172472@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	fmt.Println("Connected Mysql server")
	results, err := db.Query("SELECT * from books")
	for results.Next() {
		var book Books
		err = results.Scan(&book.ID, &book.ISBN, &book.TITLE)
		if err != nil {
			panic(err.Error())
		}
		books = append(books, book)
		fmt.Println(book)
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4040", r))
}
