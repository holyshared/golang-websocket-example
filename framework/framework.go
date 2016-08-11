package framework

import (
  "golang.org/x/net/websocket"
  "fmt"
)

func NewWebsocketServer(handlers RouteHandlers) *WebsocketServer {
  return &WebsocketServer {
    router: NewRouter(handlers),
  }
}

type WebsocketServer struct {
  router *Router
}

func (app *WebsocketServer) Handle(ws *websocket.Conn) {
  var message *SocketMessage

  err := websocket.JSON.Receive(ws, &message)

  if err != nil {
    fmt.Errorf("Unkown message error: %v", err)
    return
  }

  app.router.Emit(message.Type, ws)
}
