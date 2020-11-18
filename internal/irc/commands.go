package irc

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
)

func textParser(message twitch.PrivateMessage) (answer string) {
	switch {
	case message.Message == "!ping":
		for k, v := range message.Tags {
			fmt.Printf("%s: %s\n", k, v)
		}
		answer = "pong!"
		//}
		return
	case message.Message == "!slowmode":
		answer = "yaya function comming"
		return
	case message.Message == "!hype":
		answer = "icelanLove1 icelanCharge icelanLove1 icelanLove1 icelanSalt icelanCharge icelanCharge icelanLove1 icelanLove1"
		return
	default:
		return
	}

}
