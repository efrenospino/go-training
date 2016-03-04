package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGetTweets(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	form := url.Values{"user": {"wawandco"}, "count": {"1"}}
	request.PostForm = form
	response := httptest.NewRecorder()
	GetTweets(response, request)
	if response.Code != 200 {
		t.Fatal(err)
	}
}
