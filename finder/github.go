package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

var gitHubEndPoint = "https://api.github.com"

type GitHubUserInfo struct {
	Login    string `json:"login"`
	Avatar   string `json:"avatar_url"`
	URL      string `json:"html_url"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Repos    int    `json:"public_repos"`
	Since    string `json:"created_at"`
	Location string `json:"location"`
}

type Repository struct {
	ID          int    `jsom:"id"`
	Name        string `json:"name"`
	URL         string `json:"html_url"`
	Description string `json:"description"`
	Created     string `json:"created_at"`
	LastUpdate  string `json:"updated_at"`
	Language    string `json:"language"`
	CloneURL    string `json:"clone_url"`
}

type Data struct {
	Info  GitHubUserInfo
	Repos []Repository
}

func getUserPublicRepos(username string) []Repository {

	url1 := gitHubEndPoint + "/users/" + username + "/repos"

	req1, err := http.NewRequest("GET", url1, nil)
	CheckError(err)

	client := &http.Client{}
	response, err := client.Do(req1)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	respBody, err := ioutil.ReadAll(response.Body)

	CheckError(err)
	repos := []Repository{}
	json.Unmarshal(respBody, &repos)

	return repos

}

func FindUserInfoInGitHub(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")

	url := gitHubEndPoint + "/users/" + username

	req, err := http.NewRequest("GET", url, nil)
	CheckError(err)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	CheckError(err)

	info := GitHubUserInfo{}
	json.Unmarshal(respBody, &info)
	repos := getUserPublicRepos(username)

	t, err := template.ParseFiles("public/github.html")
	CheckError(err)

	data := Data{Info: info, Repos: repos}

	t.Execute(w, data)

}
