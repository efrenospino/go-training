package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUserInfoInGitHub(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	FindUserInfoInGitHub(response, request)
	if response.Code != 200 {
		t.Fatal(err)
	}
}
