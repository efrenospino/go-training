package main

import (
	"net/http"
	"os"
	"text/template"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func renderIndex(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("public/index.html")
	t.Execute(w, nil)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", renderIndex).Methods("GET")
	router.HandleFunc("/dribbble", FindUserInfoInDribbble).Methods("GET")
	router.HandleFunc("/github", FindUserInfoInGitHub).Methods("GET")
	router.HandleFunc("/tweets", GetTweets).Methods("POST")
	http.Handle("/", router)

	middleware := negroni.Classic()
	middleware.UseHandler(router)
	middleware.Run(":" + os.Getenv("PORT"))
}
