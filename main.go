package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// 85.Web Applications 1 - ioutil
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// /view/test
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

	// 87.template と http.ResponseWriter と http.Request
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func main() {
	// 86.Web Applications 2 -http.ListenAndServer
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	// 85.Web Applications 1 - ioutil
	/*
		p1 := &Page{Title: "test", Body: []byte("This is sample page")}
		p1.save()

		p2, _ := loadPage(p1.Title)
		fmt.Println(string(p2.Body))
	*/
}
