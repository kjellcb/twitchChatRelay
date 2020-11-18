package irc

import (
	"fmt"
	"twitchChatRelay/controllers"

	"github.com/gempir/go-twitch-irc/v2"
)

var Bot Configuration

type Configuration struct {
	Client *twitch.Client
	Cred   Credentials
	Chan   []string
	Relay  string
}

type Credentials struct {
	IsAuth bool
	User   string
	Token  string
}

type Channels struct {
	Chan string
}

func init() {
	Bot.Cred.IsAuth = true
}

func ConnectTwitch() (err error) {
	fmt.Println("Setting up a client")
	Bot.Client = twitch.NewClient(Bot.Cred.User, Bot.Cred.Token)
	fmt.Println("Connecting")
	fmt.Println("relaying chats between: " + Bot.Chan[0] + " and " + Bot.Chan[1])
	Bot.Client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		answer := textParser(message)
		if message.User.Name != "streamelements" {
			controllers.Data = append(controllers.Data, message)
		}
		if Bot.Relay == "on" {
			if message.Channel == Bot.Chan[1] {
				messageToRelay := fmt.Sprintf("C: %s - U: %s - M: %s", message.Channel, message.User.Name, message.Message)
				Bot.Client.Say(Bot.Chan[0], messageToRelay)
			}
			if message.Channel == Bot.Chan[0] {
				messageToRelay := fmt.Sprintf("C: %s - U: %s - M: %s", message.Channel, message.User.Name, message.Message)
				Bot.Client.Say(Bot.Chan[1], messageToRelay)
			}
		}
		if answer != "" {
			Bot.Client.Say(message.Channel, answer)
		}
	})

	Bot.Client.Join(Bot.Chan...)
	err = Bot.Client.Connect()

	return
}

// to be done
func relayMessage(channelIn string, channelOut string, message twitch.PrivateMessage) {
	return
}
