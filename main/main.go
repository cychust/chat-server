package main

import (
	"chat-server/internal/util/logs"
	"chat-server/models"
	"chat-server/router"
	"fmt"
	"github.com/astaxie/beego"
)

func main() {
	config := logs.Config{
		Env:  "dev",
		Path: "./log",
	}
	if err := logs.InitLogger(config); err != nil {
		return
	}
	port, _ := beego.AppConfig.Int("test::httpport")
	runPort := fmt.Sprintf("127.0.0.1:%d", port)
	router.InitWebRouter()
	models.InitOrm()

	beego.Run(runPort)
}
