package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	fmt.Println("Hello, World!")
	//var myName string = "Johnm Doe"
	//const mySecondName string = "Jane Doe"
	//
	////myThirdName := "Bob Doe"
	//fmt.Println(myName)
	//fmt.Println(mySecondName)
	app := fiber.New()
	log.Fatal(app.Listen(":4000"))
}
