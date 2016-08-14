package framework

import (
  "golang.org/x/net/websocket"
)

func NewWebsocketServer(handlers RouteHandlers, logger Logger) *WebsocketServer {
  return &WebsocketServer {
    identityNumber: 0,
    router: NewRouter(handlers),
    ConnectedContainer: NewConnectedContainer(),
    logger: logger,
  }
}

type WebsocketServer struct {
  identityNumber int
  router *Router
  logger Logger
  *ConnectedContainer
}

func (server *WebsocketServer) Handle(ws *websocket.Conn) {
  server.identityNumber++

  client := &Client {
    id: server.identityNumber,
    ws: ws,
    router: server.router,
    logger: server.logger,
    connectedServer: server,
  }
  client.Join()
  client.Read()
}
