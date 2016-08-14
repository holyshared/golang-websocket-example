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
}

func (client *Client) Read() {
  for {
    select {
      default:
        var message *SocketMessage
        err := websocket.JSON.Receive(client.ws, &message) //FIXME error

        client.logger.Info("recived message")

        if err != nil {
          log.Fatal("client.Read: %v", err.Error())
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





func NewClientContainer() *ClientContainer {
  return &ClientContainer {
    clients: map[int]*Client {},
  }
}

type ClientContainer struct {
  clients map[int]*Client
}

func (c *ClientContainer) Add(client *Client) {
  c.clients[client.id] = client
}
