package framework

func NewIdentity() *Identity {
  return &Identity {
    startValue: 0,
  }
}

type Identity struct {
  startValue int
}

func (id *Identity) nextIdentity() int {
  id.startValue++
  return id.startValue
}

func (id *Identity) currentIdentity() int {
  return id.startValue
}
