package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCalculateHandler(t *testing.T) {

	tests := []struct {
		name       string
		input      string
		wantErr    string
		wantResult float64
		wantStatus int
	}{
		{"expected input", "{\"expression\": \"2+3\"}", "", 5, 200},
		{"malformed json", "\"expression\": \"2+3\"}", "malformed json", 0, 400},
		{"shunt error", "{\"expression\": \"2+3)\"}", "unmatched closing parenthesis detected", 0, 400},
		{"eval error", "{\"expression\": \"2/0\"}", "cannot divide by zero", 0, 400},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := bytes.NewBufferString(tt.input)
			req := httptest.NewRequest(http.MethodPost, "/api/calculate", body)
			rec := httptest.NewRecorder()

			Request().ServeHTTP(rec, req)

			if rec.Code != tt.wantStatus {
				t.Fatalf("expected status %d, got %d", tt.wantStatus, rec.Code)
			}

			if tt.wantStatus != http.StatusOK {
				if !strings.Contains(rec.Body.String(), tt.wantErr) {
					t.Fatalf("expected error %s, got %s", tt.wantErr, rec.Body.String())
				} else {
					return
				}
			}

			var got struct {
				Result float64 `json:"result"`
			}
			if jsonErr := json.NewDecoder(rec.Body).Decode(&got); jsonErr != nil {
				t.Fatalf("expected answer %v, got %v", tt.wantResult, jsonErr)
				return
			}

			if got.Result != tt.wantResult {
				t.Fatalf("expected answer %v, got %v", tt.wantResult, got.Result)
			}

		})
	}
}
