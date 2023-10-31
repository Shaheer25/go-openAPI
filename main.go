package main

import (
    "log"
    "net/http"
    "strconv"

    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "github.com/go-chi/render"
)

type Book struct {
    ID     int    `json:"id"`
    Title  string `json:"title"`
    Author string `json:"author"`
}

var books = []Book{
    {1, "The Go Programming Language", "Alan A. A. Donovan, Brian W. Kernighan"},
    {2, "Designing Data-Intensive Applications", "Martin Kleppmann"},
    {3, "Code Complete", "Steve McConnell"},
}

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(render.SetContentType(render.ContentTypeJSON))

    r.Get("/books/{id}", GetBook)

    log.Println("Starting server on :3000")
    http.ListenAndServe(":3000", r)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
    bookID := chi.URLParam(r, "id")
    for _, book := range books {
        if strconv.Itoa(book.ID) == bookID {
            render.JSON(w, r, book)
            return
        }
    }
    render.Status(r, http.StatusNotFound)
}