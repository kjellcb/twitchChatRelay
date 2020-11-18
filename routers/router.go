package routers

import (
	"twitchChatRelay/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ChatController{})
}
