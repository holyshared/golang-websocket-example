package framework

type Handler func(client *Client)

func PongHandler(client *Client) {
  client.Send(Response("pong", nil))
}
