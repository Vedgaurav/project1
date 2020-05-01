package main

import (
	"fmt"
	"project1/controller"
	"project1/service/logger"
	"project1/utils"
)

func main() {

	utils.LoadEnvironmentConfiurations("env-defaults.json")
	logger.SetLogLevel(utils.LogLevel)

	router := controller.RegisterRoutes()

	fmt.Println(":" + utils.ServerPort)

	router.Run(":" + utils.ServerPort)
	//router.Run()

}
