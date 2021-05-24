package main

import (
	"fmt"

	amiClient "github.com/pooladkhay/goAmi"
)

func main() {
	client := &amiClient.Opts{
		Address:  "ASTERISK_ADDR",
		Port:     "AMI_PORT",
		Username: "AMI_USERNAME",
		Secret:   "AMI_PASSWORD",

		PingInterval:      5,
		PongTimeout:       20,
		ReconnectInterval: 2,

		EventsToListen: []string{"All"},
		EventHandler:   eventHandler,
	}

	client.Connect()
	defer client.StartListening()

	if client.Connected {
		client.SendAction("Action:PJSIPShowEndpoints")
	}
}

func eventHandler(event map[string]string) {
	if event["Event"] != "" {
		fmt.Println(event)
	}
}
