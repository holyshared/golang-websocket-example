package framework

import (
  "errors"
)

type RouteHandlers map[string]Handler

func NewRouter(handlers RouteHandlers) *Router {
  return &Router {
    routes: handlers,
  }
}

type Router struct {
  routes RouteHandlers
}

func (router *Router) Handle(route string, handler Handler) {
  router.routes[route] = handler
}

func (router *Router) Lookup(route string) (Handler, error) {
  handler, ok := router.routes[ route ]

  if !ok {
    return nil, errors.New("Unkown message type: %v")
  }

  return handler, nil
}

func (router *Router) Emit(msgType string, socket *Socket) error {
  handler, err := router.Lookup(msgType)

  if err != nil {
    return err
  }

  handler(socket)

  return nil
}
