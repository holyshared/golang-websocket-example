package framework

import (
  "golang.org/x/net/websocket"
  "log"
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
        var message *SocketMessage
        err := websocket.JSON.Receive(client.ws, &message)

        client.logger.Info("recived message")

        if err != nil {
          client.Leave()
          return
        }
        client.router.Emit(message.Type, client)
    }
  }
}

func (client *Client) Send(v interface{}) {
  err := websocket.JSON.Send(client.ws, v) //FIXME error

  if err != nil {
    log.Fatal("client.Send: %v", err.Error())
    return
  }
}

func (client *Client) Join() {
  client.logger.Infof("connect: %v", client.id)
  client.connectedServer.Add(client)
}

func (client *Client) Leave() {
  client.logger.Infof("disconnect: %v", client.id)

  defer client.ws.Close()
  client.connectedServer.Remove(client)
}
