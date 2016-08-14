package framework

type Message struct {
  Type string `json:"type"`
  Body interface{} `json:"body"`
}

func Response(msgType string, msgBody interface{}) *Message  {
  return &Message {
    Type: msgType,
    Body: msgBody,
  }
}
