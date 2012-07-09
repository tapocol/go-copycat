package hub

type Hub struct {
  Handler func(HubMessage)
  Broadcast chan HubMessage
}

func (hub Hub) Run() {
  for {
    select {
    case m := <- hub.Broadcast:
      hub.Handler(m)
    }
  }
}

type HubMessage interface {
  Format() string
}

