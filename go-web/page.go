package main 
	
import (
	"fmt"
	"io/ioutil"
	"net/http"
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

func main() {
	p1 := &Page{Title: "TestPage", }
}

func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}