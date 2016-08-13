package framework

import (
  "golang.org/x/net/websocket"
)

type Socket struct {
  ws *websocket.Conn
  logger Logger
  msg *SocketMessage
}

func (socket *Socket) Send(v interface{}) {
  err := websocket.JSON.Send(socket.ws, v)

  if err != nil {
    socket.logger.Warnf("send message error: %v", err)
    return
  }
}
