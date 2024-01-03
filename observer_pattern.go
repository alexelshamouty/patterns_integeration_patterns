package main

import (
	"observer/config"
	"observer/event_manager"
	"observer/slack_notifier"
)
import "observer/register_user"

func main() {
	register_user.Init()
	slack_notifier.Init()

	event_manager.Post_event("garbage_event")
	event_manager.Post_event(config.UserEvent)
	// Some application logic goes here to make sure that the user is registered etc, check the DB, whatever
	// Once all is fine notify via slack
	event_manager.Post_event(config.SlackEvent)
}
