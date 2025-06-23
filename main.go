package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	fmt.Println("Hello, Worlds!")
	//var myName string = "Johnm Doe"
	//const mySecondName string = "Jane Doe"
	//
	////myThirdName := "Bob Doe"
	//fmt.Println(myName)
	//fmt.Println(mySecondName)
	app := fiber.New()
	//add route
	type Todo struct {
		ID        int    `json:"id"`
		Completed bool   `json:"completed"`
		Body      string `json:"body"`
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"msg": "Hello World"})
	})

	log.Fatal(app.Listen(":4000"))
}
