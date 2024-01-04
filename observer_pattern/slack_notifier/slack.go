package slack_notifier

import "observer/event_manager"
import "fmt"
import "observer/config"

func Init() {
	//Initialise the event manager
	event_manager.Init()
	//Open Slack connection
	open_slack_connection()
	//Register handles for events
	event_manager.Subscribe(config.SlackEvent, notify_slack, notify_special_channels)
}

func notify_slack() {
	fmt.Println("Notifying slack about")
}

func notify_special_channels() {
	fmt.Println("Notifying special channels")
}
