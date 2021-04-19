package main

import (
  "net/http"
  "net/http/httptest"
  "testing"
  "github.com/stretchr/testify/assert"
)

func performRequest(router http.Handler, method, path string) *httptest.ResponseRecorder {
  request, _ := http.NewRequest(method, path, nil)
  response := httptest.NewRecorder()
  router.ServeHTTP(response, request)
  return response
}

func TestPingStatusOk(t *testing.T) {
  router := SetupRouter()

  response := performRequest(router, "GET", "/ping")

  assert.Equal(t, http.StatusOK, response.Code)
}
