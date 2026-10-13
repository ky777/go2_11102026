package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetGreet(t *testing.T) {
	req, err := http.NewRequest("GET", "http://127.0.0.1:8087/", nil)
	if err != nil {
		log.Fatal(err)
	}

	res := httptest.NewRecorder()
	if req.URL.Hostname() != "127.0.0.1" || req.URL.Port() != "8080" {
		log.Fatal("Bad host or port")
	}
	GetGreet(res, req)

	act := res.Body.String()
	exp := "<h1>Hello, I'm new web server! </h1>"
	if exp != act {
		t.Fatalf("Expected %s, but got %s", exp, act)
	}

	resCode := res.Result().StatusCode
	if resCode != 200 {
		t.Fatalf("Expected 200, but got: %d", resCode)
	}
}
