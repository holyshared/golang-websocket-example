package framework

import (
  "golang.org/x/net/websocket"
)

func NewWebsocketServer(handlers RouteHandlers, logger Logger) *WebsocketServer {
  return &WebsocketServer {
    router: NewRouter(handlers),
    logger: logger,
  }
}

type WebsocketServer struct {
  router *Router
  logger Logger
}

func (app *WebsocketServer) Handle(ws *websocket.Conn) {
  var message *SocketMessage

  err := websocket.JSON.Receive(ws, &message)

  app.logger.Infof("receive a message: %v", message)

  if err != nil {
    app.logger.Warnf("unknown message error: %v", err)
    return
  }

  socket := &Socket {
    ws: ws,
    msg: message,
    logger: app.logger,
  }

  app.router.Emit(message.Type, socket)
}
