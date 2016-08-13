package application

import (
  "github.com/holyshared/golang-websocket-example/framework"
)

func Response(type string, body interface{}) *ResponseMessage  {
  return &ResponseMessage {
    Type: type,
    Body: body,
  }
}

type ResponseMessage struct {
  Type string `json:"type"`
  Body interface{} `json:"body"`
}
