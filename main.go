package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rootiens/TwitchHub-Go/boot"
)

func main() {

	app := boot.BootApp()
	log.Fatal(app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT"))))

}
