package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	mainHandler(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("bad error code [%s]", resp.Code)
	}
}
