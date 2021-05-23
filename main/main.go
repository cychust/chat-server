package main

import (
	"chat-server/models"
	"chat-server/router"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	port, _ := beego.AppConfig.Int("test::httpport")
	runPort := fmt.Sprintf("127.0.0.1:%d", port)
	router.InitWebRouter()
	models.InitOrm()

	beego.Run(runPort)
}
