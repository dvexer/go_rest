package main

import (
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {
	req, err := http.Get("http://localhost:8000/phones/")
	if err != nil {
		t.Errorf("Error in GET(%v)", err)
	}

	if req.Status != "200 OK" {
		t.Errorf("Error status(%v)", req.Status)
	}
}
