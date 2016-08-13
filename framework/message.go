package framework

type SocketMessage struct {
  Type string `json:"type"`
  Body interface{} `json:"body"`
}

func Response(msgType string, msgBody interface{}) *ResponseMessage  {
  return &ResponseMessage {
    Type: msgType,
    Body: msgBody,
  }
}

type ResponseMessage struct {
  Type string `json:"type"`
  Body interface{} `json:"body"`
}
