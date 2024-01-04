package register_user

import (
	"fmt"
	"observer/config"
	"observer/event_manager"
)

func Init() {
	event_manager.Init()
	event_manager.Subscribe(config.UserEvent, hello, howAreYou, verifyEmail)
}

func hello() {
	fmt.Println("hello")
}

func howAreYou() {
	fmt.Println("howAreYou")
}

func verifyEmail() {
	fmt.Println("verifyEmail")
}
