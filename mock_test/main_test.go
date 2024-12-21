package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockItemService struct{}

func (m *MockItemService) GetItemByID(id int) (Item, error) {
	if id == 1 {
		return Item{ID: 1, Name: "Mock Item"}, nil
	}
	return Item{}, errors.New("item not found")
}

func TestGetItemWithInterface(t *testing.T) {

	mockService := &MockItemService{}
	router := SetupRouter(mockService)

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
			expectedBody:   `{"id":1,"name":"Mock Item"}`,
		},
		{
			name:           "Invalid ID",
			requestURL:     "/items/2",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"Item not found"}`,
		},
		{
			name:           "Invalid ID Format",
			requestURL:     "/items/abc",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid ID"}`,
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
