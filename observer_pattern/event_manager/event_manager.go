package event_manager

import "fmt"

var Subscribers map[string][]func()

func Post_event(event string) {
	_, ok := Subscribers[event]
	if ok {
		for _, callback := range Subscribers[event] {
			callback()
		}
	} else {
		fmt.Println("No subscribers for event", event)
	}
}
func Subscribe(event string, callback ...func()) {
	for _, call := range callback {
		Subscribers[event] = append(Subscribers[event], call)
	}

}

func Init() {
	if Subscribers == nil {
		Subscribers = make(map[string][]func())
	}
}
