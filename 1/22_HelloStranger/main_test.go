package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStranger(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	StrangerHandler(w, req)
	res := w.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal("server works incorrect")
	}
	if strings.TrimSpace(string(body)) != "hello stranger" {
		t.Fatal("incorrect answer", string(body), "expect 'hello stranger'")
	}

	req, err = http.NewRequest("GET", "/?name=Vasya", nil)
	if err != nil {
		t.Fatal(err)
	}
	w = httptest.NewRecorder()
	StrangerHandler(w, req)
	res = w.Result()
	defer res.Body.Close()
	body, err = io.ReadAll(res.Body)
	if err != nil {
		t.Fatal("server works incorrect")
	}
	if strings.TrimSpace(string(body)) != "hello Vasya" {
		t.Fatal("incorrect answer", string(body), "expect 'hello Vasya'")
	}
	req, err = http.NewRequest("GET", "?name=`", nil)
	if err != nil {
		t.Fatal(err)
	}
	w = httptes... File is too long to be displayed fully