package main

import (
  "net/http"
  "io"
  "golang.org/x/net/websocket"
)

func EchoServer(ws *websocket.Conn) {
  io.Copy(ws, ws);
}

func main() {
  http.Handle("/", http.FileServer( http.Dir("./public") ))
  http.Handle("/echo", websocket.Handler(EchoServer));

  err := http.ListenAndServe(":3000", nil);

  if err != nil {
    panic("ListenAndServe: " + err.Error())
  }
}
