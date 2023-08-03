package main

import (
	"alerts/modules/alerts"
	"fmt"
)

func main() {
	env := LoadEnv()

	alertModule := alerts.CreateModule()
	err := alertModule.Initialize()
	if err != nil {
		panic("Error while initializing alert module!")
	}

	fmt.Println("Brick by brick...")

	StartServer(env.Port)

}
