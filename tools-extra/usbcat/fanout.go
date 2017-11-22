package main

type Client chan Message

type ActionType string

const (
	ATSubscribe   ActionType = "sub"
	ATUnsubscribe ActionType = "unsub"
)

type Action struct {
	Type ActionType
	Data interface{}
}

type DynamicFanOut struct {
	source  <-chan Message
	action  chan Action
	clients []Client
}

func runDFO(src <-chan Message) *DynamicFanOut {
	dfo := &DynamicFanOut{
		source:  src,
		clients: []Client{},
		action:  make(chan Action),
	}
	go dfo.run()
	return dfo
}

func (dfo *DynamicFanOut) run() {
	for {
		select {
		case msg := <-dfo.source:
			for _, chn := range dfo.clients {
				chn <- msg
			}
		case action := <-dfo.action:
			switch action.Type {
			case ATSubscribe:
				newchan := action.Data.(Client)
				dfo.clients = append(dfo.clients, newchan)
			case ATUnsubscribe:
				oldchan := action.Data.(Client)
				for i, chn := range dfo.clients {
					if chn == oldchan {
						dfo.clients[i] = dfo.clients[len(dfo.clients)-1]
						dfo.clients = dfo.clients[:len(dfo.clients)-1]
						break
					}
				}
			}
		}
	}
}

func (dfo *DynamicFanOut) Subscribe() Client {
	client := make(Client)
	dfo.action <- Action{
		Type: ATSubscribe,
		Data: client,
	}
	return client
}

func (dfo *DynamicFanOut) Unsubscribe(client Client) {
	dfo.action <- Action{
		Type: ATUnsubscribe,
		Data: client,
	}
}
