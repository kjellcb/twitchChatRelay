package routers

import (
	"twitchchatrelay/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.ChatController{})
}
