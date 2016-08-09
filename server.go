package main

import (
  "net/http"
//  "io"
  //"encoding/json"
  "golang.org/x/net/websocket"
)
/*
type Message struct {
  Message string `json:"message"`
}

func (m *Message) WriteTo(w io.Writer) (n int64, err error) {
  bytes, err := json.Marshal(m)
  if err != nil {
    return 0, err
  }
  length, werr := w.Write(bytes)

  return int64(length), werr
}
*/
/*
func (m *Message) MarshalJSON() ([]byte, error) {
  return json.Marshal(m)
}
*/
/*
func ResponseMessage(message string) *Message {
  return &Message {
    Message: message,
  }
}
*/
func EchoServer(ws *websocket.Conn) {
  message := ResponseMessage("ok!!")
  message.WriteTo(ws)
}

func main() {
  http.Handle("/", http.FileServer( http.Dir("./public") ))
  http.Handle("/echo", websocket.Handler(EchoServer));

  err := http.ListenAndServe(":3000", nil);

  if err != nil {
    panic("ListenAndServe: " + err.Error())
  }
}
