package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	jsonData, err := os.Open("./gopher.json")
	if err != nil {
		panic(err)
	}
	defer jsonData.Close()
	byteValue, _ := ioutil.ReadAll(jsonData)

	var story map[string]Arc
	err = json.Unmarshal(byteValue, &story)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.ParseFiles("story_page.html")
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", cyoaHandler{story, tmpl})
}

type cyoaHandler struct {
	story map[string]Arc
	tmpl  *template.Template
}

func (th cyoaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = path[1:]
	if path == "" {
		path = "intro"
	}

	if arc, ok := th.story[path]; ok {
		th.tmpl.Execute(w, arc)
	} else {
		w.Write([]byte("how did you get here?"))
	}
}

type Arc struct {
	Title   string
	Story   []string
	Options []Options
}

type Options struct {
	Text string
	Arc  string
}
