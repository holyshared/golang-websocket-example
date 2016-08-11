package framework

import (
  "golang.org/x/net/websocket"
  "fmt"
)

type Socket struct {
  ws *websocket.Conn
  msg *SocketMessage
}

func (socket *Socket) Send(v interface{}) {
  err := websocket.JSON.Send(socket.ws, v)

  if err != nil {
    fmt.Errorf("send message error: %v", err)
    return
  }
}
