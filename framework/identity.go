package framework

func NewIdentity() *Identity {
  return &Identity {
    startValue: 0,
  }
}

type Identity struct {
  startValue int
}

func (id *Identity) next() int {
  id.startValue++
  return id.startValue
}

func (id *Identity) current() int {
  return id.startValue
}
