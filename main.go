package main

import (
	"fizzbuzz/controllers"
	"fmt"
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	fmt.Println("-- Server Started --")

	app := fiber.New()
	controller := new(controllers.FizzControl)

	app.Post("/fizzbuzz", controller.HandleFizzBuzz)
	app.Get("/stats", controller.GetStats)

	log.Fatal(app.Listen(":8080"))
}
