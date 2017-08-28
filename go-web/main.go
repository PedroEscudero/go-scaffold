package main 
	
import (
	"fmt"
	"io/ioutil"

)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.title + ".txt"
	return ioutil.WriteFile(filename, p.body, 0600)
}

func loadPage (title string) *Page {
	filename := title + ".txt"
	body, _ := ioutil.ReadFile(filename)
}

func loadPage(title string) *Page {
    filename := title + ".txt"
    if err != nil {
    	return nil, err
    }
    body, _ := ioutil.ReadFile(filename)
    return &Page{Title: title, Body: body}, nil
}