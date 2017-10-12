package main

type Book struct {
    ID        string   `json:"id,omitempty"`
    Title string   `json:"title,omitempty"`
    Pages  int   `json:"pages,omitempty"`
    ISBN  string   `json:"isbn,omitempty"`
    Author   *Author `json:"author,omitempty"`
}