package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

var testingCall = false

const (
	endPointURL = "https://api.flickr.com/services/rest/"
	apiKey      = "0a97e077103df1244029fdab72379f39"
	method      = "flickr.photos.search"
	media       = "jpg"
	format      = "json"
)

//JSONObject is a JSON type
type JSONObject struct {
	Photos PhotosType
	Stat   string
}

//PhotosType is a Photos type, contains all photos
type PhotosType struct {
	Page       int
	Pages      int
	PerPage    int
	Total      string
	PhotosList []PhotoType `json:"photo"`
}

//PhotoType is a Photo type, contains information about it
type PhotoType struct {
	ID     string
	Secret string
	Server string
	Farm   int
	Title  string
}

type PhotoXML struct {
	ID     string `xml:"id,attr"`
	Secret string `xml:"secret,attr"`
	Server string `xml:"server,attr"`
	Farm   int    `xml:"farm,attr"`
	Title  string `xml:"title,attr"`
}

type PhotosListXML struct {
	PhotosList []PhotoXML `xml:"photo"`
}

type XMLObject struct {
	XMLName xml.Name      `xml:"rsp"`
	Status  string        `xml:"stat,attr"`
	Photos  PhotosListXML `xml:"photos"`
}

var flickrResponseBody []byte

func formatRequestTo(text string) string {
	request := endPointURL + "?api_key=" + apiKey + "&method=" + method + "&text=" + text + "&media=" + media + "&format=" + format + "&per_page=50&tags=" + text
	return request
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func decodeXML(body []byte) (xmlType XMLObject) {
	err := xml.Unmarshal(body, &xmlType)
	checkError(err)
	return xmlType
}

func decodeJSON(body []byte) (jsontype JSONObject) {
	bodyStr := strings.Replace(string(body), "jsonFlickrApi(", "", 1)
	bodyStr = strings.Replace(bodyStr, "})", "}", 1)
	err := json.Unmarshal([]byte(bodyStr), &jsontype)
	checkError(err)
	return jsontype
}

func renderForm(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("public/search.html")
	t.Execute(w, nil)
}

func searchTo(w http.ResponseWriter, r *http.Request) {
	if !testingCall {
		r.ParseForm()
		url := strings.Replace(formatRequestTo(r.FormValue("text")), " ", "%20", -1)
		req, err := http.NewRequest("POST", url, nil)
		//req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		flickrResponseBody, _ = ioutil.ReadAll(resp.Body)
	}

	t, _ := template.ParseFiles("public/results.html")

	fmt.Println("Format:", format)
	if format == "rest" {
		decoded := decodeXML(flickrResponseBody)
		t.Execute(w, decoded.Photos.PhotosList)
	} else if format == "json" {
		decoded := decodeJSON(flickrResponseBody)
		t.Execute(w, decoded.Photos.PhotosList)
	}

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", renderForm).Methods("GET")
	router.HandleFunc("/search", searchTo).Methods("POST")
	http.Handle("/", router)

	n := negroni.Classic()
	n.UseHandler(router)

	port := os.Getenv("PORT")

	n.Run(":" + port)
}
