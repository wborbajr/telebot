package main

import (
	"fmt"
	routes "github.com/wborbajr/telebot/routes"
)

func main() {
	fmt.Println("TeleBOTServer v1.0.0")
	routes.StartGin()
}