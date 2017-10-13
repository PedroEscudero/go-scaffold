package main

import (
    "encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"strings"
)

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var books []Book
var	router = mux.NewRouter()

var db *sql.DB 
var err error

type Author struct {
    Name  string `json:"name,omitempty"`
    Nacionality string `json:"nacionality,omitempty"`
}

func main() {
    setRoutes()
    seedBooks()
    log.Fatal(http.ListenAndServe(":8080", router))
}

func seedBooks() {
	books = append(books, Book{ID: "1", Title: "Bible", Pages: 1000, ISBN: "weweuruiererer-4", Author: &Author{Name: "God", Nacionality: "All"}})
	books = append(books, Book{ID: "2", Title: "It", Pages: 500, ISBN: "weweuruiererer-5", Author: &Author{Name: "Sthepen King", Nacionality: "USA"}})
	books = append(books, Book{ID: "3", Title: "Zombikindergarten", Pages: 250, ISBN: "miimimie3", Author: &Author{Name: "Pedro Escudero", Nacionality: "España"}})
}

func setRoutes(){
	setRoute("book")
}

func setRoute(route string){
	router.HandleFunc("/{route}", GetBooks).Methods("GET")
	router.HandleFunc("/{route}/{id}", GetBook).Methods("GET")
	router.HandleFunc("/{route}/{id}", CreateBook).Methods("POST")
    router.HandleFunc("/{route}/{id}", Delete).Methods("DELETE")
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	model := strings.Replace(r.URL.Path, "/", "", -1)
	if model == "books"{
		json.NewEncoder(w).Encode(books)
	}else{
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
    var book Book

 	db, err = sql.Open("mysql", "roort:@/test")
    if err != nil {
        panic(err.Error())    
    }

    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }

    params := mux.Vars(r)

    db.QueryRow("SELECT id, title, ISBN, pages FROM books WHERE id=?", params["id"]).Scan(&book.ID , &book.Title , &book.ISBN, &book.Pages)
    json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
    var book Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    book.ID = params["id"]
    books = append(books, book)
    json.NewEncoder(w).Encode(books)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range books {
        if item.ID == params["id"] {
            books = append(books[:index], books[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(books)
    }
}
