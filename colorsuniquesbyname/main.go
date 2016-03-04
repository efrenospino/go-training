package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/efrenospino/generatecolor"
)

func generateColorHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("public/index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		value := generatecolor.GenerateValuesByLettersInName(r.Form["name"][0])
		t, _ := template.ParseFiles("public/color.html")
		t.Execute(w, generatecolor.Person{Name: strings.ToUpper(r.Form["name"][0]), Color: value})
	}
}

func main() {
	http.HandleFunc("/", generateColorHandle)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
