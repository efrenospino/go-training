package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

const (
	dribbbleClientAccessToken = "ea16118a37c0644ed97a422191c46e44013ad3ae333b33c4eab9b6d9507503b9"
	endPoint                  = "https://api.dribbble.com/v1"
)

type DribbbleUserShot struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Images      struct {
		Normal string `json:"normal"`
	} `json:"images"`
	Views int    `json:"views_count"`
	Likes int    `json:"likes_count"`
	Date  string `json:"created_at"`
}

type DribbbleUserInfoResponse struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	URL      string `json:"html_url"`
	Avatar   string `json:"avatar_url"`
	Bio      string `json:"bio"`
	Location string `json:"location"`
	Links    struct {
		Website string `json:"web"`
		Twitter string `json:"twitter"`
	} `json:"links"`
	Since string `json:"created_at"`
}

type TemplateData struct {
	Info  DribbbleUserInfoResponse
	Shots []DribbbleUserShot
}

func FindUserShotsInDribbble(username string) []DribbbleUserShot {

	url1 := endPoint + "/users/" + username + "/shots?access_token=" + dribbbleClientAccessToken

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
	shots := []DribbbleUserShot{}
	json.Unmarshal(respBody, &shots)

	return shots

}

func FindUserInfoInDribbble(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")

	url1 := endPoint + "/users/" + username + "?access_token=" + dribbbleClientAccessToken

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
	info := DribbbleUserInfoResponse{}
	json.Unmarshal([]byte(respBody), &info)
	shots := FindUserShotsInDribbble(username)

	t, _ := template.ParseFiles("public/dribbble.html")
	data := TemplateData{Info: info, Shots: shots}

	t.Execute(w, data)

}
