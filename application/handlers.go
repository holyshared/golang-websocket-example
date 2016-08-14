package application

import (
  "github.com/holyshared/golang-websocket-example/framework"
)

func Test(client *framework.Client) {
  client.Send(framework.Response("test", "test"))
}
