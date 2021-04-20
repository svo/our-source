package main

import (
  "encoding/json"
  "net/http"
  "net/http/httptest"
  "testing"

  "github.com/gin-gonic/gin"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/suite"
)

type MainSuite struct {
  suite.Suite
  router http.Handler
}

func performRequest(router http.Handler, method, path string) *httptest.ResponseRecorder {
  request, _ := http.NewRequest(method, path, nil)
  response := httptest.NewRecorder()
  router.ServeHTTP(response, request)
  return response
}

func (suite *MainSuite) SetupTest() {
  suite.router = setupRouter()
}

func (suite *MainSuite) TestPingStatusOk() {
  response := performRequest(suite.router, "GET", "/ping")

  assert.Equal(suite.T(), http.StatusOK, response.Code)
}

func (suite *MainSuite) TestPingMessageBody() {
  body := gin.H{
    "message": "pong",
  }

  response := performRequest(suite.router, "GET", "/ping")

  var response_body map[string]string
  err := json.Unmarshal(response.Body.Bytes(), &response_body)
  value, exists := response_body["message"]

  assert.Nil(suite.T(), err)
  assert.True(suite.T(), exists)
  assert.Equal(suite.T(), body["message"], value)
}

func TestMainSuite(t *testing.T) {
  suite.Run(t, new(MainSuite))
}
