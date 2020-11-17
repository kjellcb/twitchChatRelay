package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"twitchChatRelay/internal/irc"

	_ "twitchchatrelay/routers"

	"github.com/astaxie/beego"
)

func init() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()
}

func main() {
	fmt.Println("Starting Webserver at :11000")
	go func() {
		beego.Run()
	}()
	fmt.Println("Starting Twitch Bot")
	err := irc.ConnectTwitch()
	fmt.Println(err)
}
