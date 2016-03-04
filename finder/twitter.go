package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mrjones/oauth"
)

type User struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

type Tweet struct {
	Date string `json:"created_at"`
	Text string `json:"text"`
	User User   `json:"user"`
}

//Constants for twitter API
const (
	RequestURL        = "https://api.twitter.com/1.1/statuses/user_timeline.json"
	ConsumerKey       = "xNNAL1902gzeml4E3IO8NTR5X"
	ConsumerSecret    = "tfwPvvdHjWZB5CAehwCChxtFt2KgScMYte4EmsqBafEbyy3VVV"
	AccessToken       = "868832618-QnG1pxCkjFpsSIMl8HZAOg3J8FnsCFJVA9qFQi6t"
	AccessTokenSecret = "KlHYyv95yeXixuMgmsa3NByXxC7Edfxmayp7FPeQreDfr"
)

var params = make(map[string]string)
var tweets = make([]Tweet, 0)

func GetTweets(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	params["screen_name"] = r.FormValue("user")
	params["count"] = r.FormValue("count")
	consumer := oauth.NewConsumer(ConsumerKey, ConsumerSecret, oauth.ServiceProvider{})
	token := &oauth.AccessToken{Token: AccessToken, Secret: AccessTokenSecret}
	response, err := consumer.Get(RequestURL, params, token)
	CheckError(err)

	defer response.Body.Close()
	respBody, err := ioutil.ReadAll(response.Body)
	CheckError(err)
	json.Unmarshal([]byte(respBody), &tweets)

	t, _ := template.ParseFiles("public/tweets.html")
	t.Execute(w, tweets)
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}
