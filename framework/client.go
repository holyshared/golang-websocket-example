package framework

import (
  "golang.org/x/net/websocket"
)

func NewClient(id int, ws *websocket.Conn, router *Router, logger Logger) *Client {
  return &Client {
    id: id,
    ws: ws,
    router: router,
    logger: logger,
  }
}

type Client struct {
  id int
  ws *websocket.Conn
  router *Router
  logger Logger
  connectedServer ClientContainer
}

func (client *Client) Read() {
  for {
    select {
      default:
        var message *Message
        err := websocket.JSON.Receive(client.ws, &message)

        if err != nil {
          client.LeaveFromServer()
          return
        }
        client.logger.Infof("recived message: %v %v", client.id, message)
        client.router.Emit(message.Type, client)
    }
  }
}

func (client *Client) Send(v interface{}) {
  err := websocket.JSON.Send(client.ws, v) 

  if err != nil {
    client.LeaveFromServer()
    return
  }
  client.logger.Infof("send message: %v %v", client.id, v)
}

func (client *Client) JoinToServer() {
  client.logger.Infof("connect: %v", client.id)
  client.connectedServer.Add(client)
}

func (client *Client) LeaveFromServer() {
  client.logger.Infof("disconnect: %v", client.id)

  defer client.ws.Close()
  client.connectedServer.Remove(client)
}
