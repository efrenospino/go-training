package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestSearchToHandler_GET(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	generateColorHandle(response, request)
	if response.Code != 200 {
		t.Fatal(response.Code)
		return
	}
}

func TestSearchToHandler_POST(t *testing.T) {
	request, err := http.NewRequest("POST", "/generate-color", nil)
	if err != nil {
		t.Fatal(err)
	}

	form := url.Values{"name": {"Efren"}}
	request.PostForm = form
	request.ParseForm()

	response := httptest.NewRecorder()
	generateColorHandle(response, request)
	fmt.Println(response)
	if response.Code != 200 {
		t.Fatal(response.Code)
		return
	}
}
