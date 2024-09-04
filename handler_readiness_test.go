package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	tests := []struct {
		name         string
		code         int
		msg          string
		expectedBody string
		expectLog    bool
	}{
		{
			name:         "4XX Error",
			code:         http.StatusBadRequest,
			msg:          "Bad request",
			expectedBody: `{"error":"Bad request"}`,
			expectLog:    false,
		},
		{
			name:         "5XX Error",
			code:         http.StatusInternalServerError,
			msg:          "Internal server error",
			expectedBody: `{"error":"Internal server error"}`,
			expectLog:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture log output
			var logOutput bytes.Buffer
			log.SetOutput(&logOutput)

			rr := httptest.NewRecorder()

			respondWithError(rr, tt.code, tt.msg)

			body, err := io.ReadAll(rr.Body)
			if err != nil {
				t.Fatal(err)
			}

			if strings.TrimSpace(string(body)) != tt.expectedBody {
				t.Errorf("Expected body %s, got %s", tt.expectedBody, body)
			}

			if rr.Code != tt.code {
				t.Errorf("Expected status code %d, got %d", tt.code, rr.Code)
			}

			if tt.expectLog {
				if !strings.Contains(logOutput.String(), "Responding with 5XX err:") {
					t.Errorf("Expected log output, got none")
				}
			} else if strings.Contains(logOutput.String(), "Responding with 5XX err:") {
				t.Errorf("Expected no log output, but got some")
			}
		})
	}
}

func TestHandlerReadiness(t *testing.T) {
	rr := httptest.NewRecorder()

	req := httptest.NewRequest("GET", "/readiness", nil)

	handlerReadiness(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", rr.Code)
	}

	expectedBody := `{}`

	body, err := io.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}

	if strings.TrimSpace(string(body)) != expectedBody {
		t.Errorf("Expected body %s, got %s", expectedBody, body)
	}
}
