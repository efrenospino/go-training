package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindUserInfoInDribbble(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	FindUserInfoInDribbble(response, request)
	if response.Code != 200 {
		t.Fatal(err)
	}
}
