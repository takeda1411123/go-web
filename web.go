package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

var cs *sessions.CookieStore = sessions.NewCookieStore([]byte("secret-key-12345"))

func notemp() *template.Template {
	src := "<html><body><h1>NO Template.</h1></body></html>"
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

// get target Temlate.
func page(fname string) *template.Template {
	tmps, _ := template.ParseFiles("templates/"+fname+".html",
		"templates/head.html", "templates/foot.html")
	return tmps
}

// index handler
func index(w http.ResponseWriter, rq *http.Request) {
	item := struct {
		Template string
		Message  string
		Title    string
	}{
		Message:  "This is Top page.",
		Title:    "Index",
		Template: "index",
	}
	er := page("index").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

// hello handler
func hello(w http.ResponseWriter, rq *http.Request) {
	data := []string{
		"one", "two", "three",
	}
	item := struct {
		Title string
		Data  []string
	}{
		Title: "Hello",
		Data:  data,
	}
	er := page("hello").Execute(w, item)
	if er != nil {
		log.Fatal(er)
	}
}

func main() {

	// index handling
	http.HandleFunc("/", func(w http.ResponseWriter, rq *http.Request) {
		index(w, rq)
	})
	// hello handling
	http.HandleFunc("/hello", func(w http.ResponseWriter, rq *http.Request) {
		hello(w, rq)
	})

	if err := http.ListenAndServe("", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
