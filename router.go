package main

import (
  "encoding/json"
  "golang.org/x/net/websocket"
  "fmt"
)

/*
func NewRouter(ws *websocket.Conn) *Router {
  return &Router {
  ws: ws,
  }
}
*/

type SocketMessage interface {
  MessageType() string
}

type ReceiveSocketMessage struct {
  Type string `json:"type"`
}

func (msg *ReceiveSocketMessage) MessageType() string {
  return msg.Type
}

type Socket struct {
  message SocketMessage
  ws *websocket.Conn
}

type Router struct {
  routes map[string]func(socket *Socket)
}

func (router *Router) Handle(route string, handler func(socket *Socket)) {
  router.routes[route] = handler
}

func (router *Router) Dispatch(ws *websocket.Conn) {
  var message []byte
  var receiveMessage *ReceiveSocketMessage
  ws.Read(message)

  err := json.Unmarshal(message, receiveMessage)

  if err != nil {
    fmt.Errorf("Unkown message error: %v", err)
    return
  }

  handler, ok := router.routes[ receiveMessage.MessageType() ]

  if !ok {
    return
  }

  socket := &Socket {
    message: receiveMessage,
    ws: ws,
  }
  handler(socket)
}
