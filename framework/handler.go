package framework

type Handler func(socket *Socket)

func PongHandler(socket *Socket) {
  socket.Send(Response("pong", nil))
}
