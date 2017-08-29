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

func viewHandle (w http.ResponseWriter, r *http.request){
	title := r.URL.Path[len("/view/")]
	  p, _ := loadPage(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.ListenAndServe(":8080", nil)
}

