package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleErr(t *testing.T) {
	rr := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/error", nil)

	handleErr(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}

	expectedBody := `{"error":"Bad Request"}`

	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(string(body)) != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, body)
	}
}
