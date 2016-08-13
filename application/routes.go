package application

import (
  "github.com/holyshared/golang-websocket-example/framework"
)

var (
  Routes = map[string]framework.Handler {
    "test": Test,
    "ping": framework.PongHandler,
  }
)
