package main

import (
	"github.com/efrenospino/firstRaspao/config"
	_ "github.com/efrenospino/firstRaspao/controllers"
	raspao "github.com/wawandco/raspao/base"
)

func main() {
	config.DefineRoutes(&raspao.Router{})
	config.SharedApplication = raspao.NewApp()
	config.SharedApplication.Start()

	defer config.SharedApplication.Database.Close()
}
