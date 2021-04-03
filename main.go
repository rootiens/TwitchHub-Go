package main

import (
	"TwitchHub/boot"
	"fmt"
	"log"
	"os"
)

func main() {

	app := boot.BootApp()
	log.Fatal(app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT"))))

}
