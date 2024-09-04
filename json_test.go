package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRespondWithJSON_Success(t *testing.T) {
	rr := httptest.NewRecorder()
	payload := map[string]string{"message": "success"}

	respondWithJSON(rr, http.StatusOK, payload)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, status)
	}

	expected := `{"message":"success"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("expected body %s, got %s", expected, rr.Body.String())
	}
}

func TestRespondWithJSON_InternalServerError(t *testing.T) {
	rr := httptest.NewRecorder()
	payload := make(chan int)

	respondWithJSON(rr, http.StatusOK, payload)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, status)
	}

	expected := "Internal Server Error\n"
	if rr.Body.String() != expected {
		t.Errorf("expected body %s, got %s", expected, rr.Body.String())
	}
}
