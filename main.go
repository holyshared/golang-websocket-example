package main

import (
  "net/http"
  "golang.org/x/net/websocket"
  "github.com/holyshared/golang-websocket-example/application"
  "github.com/holyshared/golang-websocket-example/framework"
)

func main() {
  server := framework.NewWebsocketServer(
    application.Routes,
    application.Logger,
  )

  http.Handle("/", http.FileServer( http.Dir("./public") ))
  http.Handle("/echo", websocket.Handler(server.Handle))

  err := http.ListenAndServe(":3000", nil);

  if err != nil {
    panic("ListenAndServe: " + err.Error())
  }
}
