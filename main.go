package main

import (
  "log"

  "github.com/svo/our-source/golang/server"
)

func main() {
  err := server.Init(server.GinRouterFactory{}).Run()

  if err != nil {
    log.Fatal(err)
  }
}
