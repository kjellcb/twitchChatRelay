package controllers

import (
	"github.com/astaxie/beego"
	"github.com/gempir/go-twitch-irc/v2"
)

type ChatController struct {
	beego.Controller
}

var Data [](twitch.PrivateMessage)

func (ctx *ChatController) Get() {

	ctx.Data["messages"] = &Data
	ctx.TplName = "index.html"
}
