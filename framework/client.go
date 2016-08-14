package framework

import (
  "golang.org/x/net/websocket"
)

func NewClient(id int, connection *websocket.Conn, router *Router, logger Logger) *Client {
  return &Client {
    id: id,
    connection: connection,
    router: router,
    logger: logger,
  }
}

type Client struct {
  id int
  connection *websocket.Conn
  router *Router
  logger Logger
  connectedServer ClientContainer
}

func (client *Client) Read() {
  for {
    select {
      default:
        var message *Message
        err := websocket.JSON.Receive(client.connection, &message)

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
  err := websocket.JSON.Send(client.connection, v) 

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

  defer client.connection.Close()
  client.connectedServer.Remove(client)
}
