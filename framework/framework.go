package framework

import (
  "golang.org/x/net/websocket"
)

func NewWebsocketServer(handlers RouteHandlers, logger Logger) *WebsocketServer {
  return &WebsocketServer {
    identityNumber: 0,
    router: NewRouter(handlers),
    ClientContainer: NewClientContainer(),
    logger: logger,
  }
}

type WebsocketServer struct {
  identityNumber int
  router *Router
  logger Logger
  *ClientContainer
}

func (server *WebsocketServer) Handle(ws *websocket.Conn) {
  server.identityNumber++
  server.logger.Infof("connected: %v", server.identityNumber)

  client := &Client {
    id: server.identityNumber,
    ws: ws,
    router: server.router,
    logger: server.logger,
  }
  server.Add(client)
  client.Read()
}
