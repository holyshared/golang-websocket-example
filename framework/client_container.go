package framework

func NewConnectedContainer() *ConnectedContainer {
  return &ConnectedContainer {
    clients: map[int]*Client {},
  }
}

type ClientContainer interface {
  add(client *Client)
  remove(client *Client)
}

type ConnectedContainer struct {
  clients map[int]*Client
}

func (c *ConnectedContainer) add(client *Client) {
  c.clients[client.id] = client
}

func (c *ConnectedContainer) remove(client *Client) {
  delete(c.clients, client.id)
}
