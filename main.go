package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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

// 89.templateのキャッシュとハンドラー
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// t, _ := template.ParseFiles(tmpl + ".html")
	// t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	// /view/test
	// title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

	// 87.Web Applications 3 - template と http.ResponseWriter と http.Request
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// 88.Web Applications 4 - http.Redirect
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	// title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var validPath = regexp.MustCompile("^/(edit|save|view|)/([a-zA-Z0-0]+)$")

func makeHundler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func main() {
	// 86.Web Applications 2 -http.ListenAndServer
	http.HandleFunc("/view/", makeHundler(viewHandler))
	http.HandleFunc("/edit/", makeHundler(editHandler))
	http.HandleFunc("/save/", makeHundler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))

	// 85.Web Applications 1 - ioutil
	/*
		p1 := &Page{Title: "test", Body: []byte("This is sample page")}
		p1.save()

		p2, _ := loadPage(p1.Title)
		fmt.Println(string(p2.Body))
	*/
}
