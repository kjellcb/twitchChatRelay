package irc

import (
	"fmt"
	"twitchchatrelay/controllers"

	"github.com/gempir/go-twitch-irc"
)

var Bot Configuration

type Configuration struct {
	Client *twitch.Client
	Cred   Credentials
	Chan   []string
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
	Bot.Cred.User = "mancert"
	Bot.Cred.Token = "oauth:9y7q7c9xffmpgs63ue1e7klcv4z4wl"
	Bot.Chan = []string{
		"mancert",
		"icelandicice",
	}
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

		if message.Channel == Bot.Chan[1] {
			messageToRelay := fmt.Sprintf("C: %s - U: %s - M: %s", message.Channel, message.User.Name, message.Message)
			Bot.Client.Say(Bot.Chan[0], messageToRelay)
		}
		if message.Channel == Bot.Chan[0] {
			messageToRelay := fmt.Sprintf("C: %s - U: %s - M: %s", message.Channel, message.User.Name, message.Message)
			Bot.Client.Say(Bot.Chan[1], messageToRelay)
		}
		if answer != "" {
			Bot.Client.Say(message.Channel, answer)
		}
	})

	Bot.Client.Join(Bot.Chan...)
	err = Bot.Client.Connect()

	return
}

func relayMessage(channelIn string, channelOut string, message twitch.PrivateMessage) {

}
