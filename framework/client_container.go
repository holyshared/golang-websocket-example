package framework

func NewConnectedContainer() *ConnectedContainer {
  return &ConnectedContainer {
    clients: map[int]*Client {},
  }
}

type ClientContainer interface {
  Add(client *Client)
  Remove(client *Client)
}

type ConnectedContainer struct {
  clients map[int]*Client
}

func (c *ConnectedContainer) Add(client *Client) {
  c.clients[client.id] = client
}

func (c *ConnectedContainer) Remove(client *Client) {
  delete(c.clients, client.id)
}
