package application

import (
  "github.com/holyshared/golang-websocket-example/framework"
)

func Test(socket *framework.Socket) {
  socket.Send("ok")
}
