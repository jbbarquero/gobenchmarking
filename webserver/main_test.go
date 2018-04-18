package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDobleHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "localhost:8088/double?v=2", nil)
	if err != nil {
		t.Fatalf("could not create request %v", err)
	}

	rec := httptest.NewRecorder()

	doubleHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	log.Println(res.Status)
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
	if err != nil {
		t.Fatalf("expected an integer; got %s", b)
	}

	if d != 4 {
		t.Fatalf("expected double to be 4; got %v", d)
	}

}
