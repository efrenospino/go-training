package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/efrenospino/parser/jsontoxml"
	"github.com/efrenospino/parser/xmltojson"
	"github.com/gorilla/mux"
)

func renderIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("public/index.html")
	t.Execute(w, nil)
}

func renderToJSONForm(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("public/tojson.html")
	t.Execute(w, nil)
}

func renderToXMLForm(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("public/toxml.html")
	t.Execute(w, nil)
}

func parseXMLToJSON(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	value := xmltojson.XMLToJSON(r.FormValue("xml-to-json-input"))
	t, _ := template.ParseFiles("public/tojson.html")
	t.Execute(w, value)
}

func parseJSONToXML(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	value := jsontoxml.JSONToXML(r.FormValue("json-to-xml-input"))
	t, _ := template.ParseFiles("public/toxml.html")
	t.Execute(w, value)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/", renderIndex).Methods("GET")

	router.HandleFunc("/xml-to-json", renderToJSONForm).Methods("GET")
	router.HandleFunc("/xml-to-json", parseXMLToJSON).Methods("POST")

	router.HandleFunc("/json-to-xml", renderToXMLForm).Methods("GET")
	router.HandleFunc("/json-to-xml", parseJSONToXML).Methods("POST")

	http.Handle("/", router)

	n := negroni.Classic()
	n.UseHandler(router)

	port := os.Getenv("PORT")

	n.Run(":" + port)
}
