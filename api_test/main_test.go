package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetItem(t *testing.T) {

	router := SetupRouter()

	tests := []struct {
		name           string
		requestURL     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid ID",
			requestURL:     "/items/1",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"id":1,"name":"Item One"}`,
		},
		{
			name:           "Invalid ID",
			requestURL:     "/items/2",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"Item not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req, _ := http.NewRequest(http.MethodGet, tt.requestURL, nil)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedStatus, resp.Code)
			assert.JSONEq(t, tt.expectedBody, resp.Body.String())
		})
	}
}
