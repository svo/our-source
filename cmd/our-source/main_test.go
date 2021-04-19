package main

import (
  "encoding/json"
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/gin-gonic/gin"
  "github.com/stretchr/testify/assert"
)

func performRequest(router http.Handler, method, path string) *httptest.ResponseRecorder {
  request, _ := http.NewRequest(method, path, nil)
  response := httptest.NewRecorder()
  router.ServeHTTP(response, request)
  return response
}

func TestPingStatusOk(t *testing.T) {
  router := setupRouter()

  response := performRequest(router, "GET", "/ping")

  assert.Equal(t, http.StatusOK, response.Code)
}

func TestPingMessageBody(t *testing.T) {
  router := setupRouter()
  body := gin.H{
    "message": "pong",
  }

  response := performRequest(router, "GET", "/ping")

  var response_body map[string]string
  err := json.Unmarshal(response.Body.Bytes(), &response_body)
  value, exists := response_body["message"]

  assert.Nil(t, err)
  assert.True(t, exists)
  assert.Equal(t, body["message"], value)
}
