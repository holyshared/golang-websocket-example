package framework

import (
  "golang.org/x/net/websocket"
)

func NewWebsocketServer(handlers RouteHandlers, logger Logger) *WebsocketServer {
  return &WebsocketServer {
    Identity: NewIdentity(),
    router: NewRouter(handlers),
    ConnectedContainer: NewConnectedContainer(),
    logger: logger,
  }
}

type WebsocketServer struct {
  router *Router
  logger Logger
  *ConnectedContainer
  *Identity
}

func (server *WebsocketServer) Handle(ws *websocket.Conn) {
  server.nextIdentity()

  client := &Client {
    id: server.currentIdentity(),
    ws: ws,
    router: server.router,
    logger: server.logger,
    connectedServer: server,
  }
  client.JoinToServer()
  client.Read()
}
