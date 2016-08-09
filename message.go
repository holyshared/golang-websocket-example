package main

import (
  "io"
  "encoding/json"
)

func ResponseMessage(message string) *Message {
  return &Message {
    Message: message,
  }
}

type Message struct {
  Message string `json:"message"`
}

func (m *Message) WriteTo(w io.Writer) (n int64, err error) {
  bytes, err := json.Marshal(m)
  if err != nil {
    return int64(len(bytes)), err
  }
  length, werr := w.Write(bytes)

  return int64(length), werr
}
