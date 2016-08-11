package framework

type SocketMessage struct {
  Type string `json:"type"`
  Body interface{} `json:"body"`
}
