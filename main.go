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

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"msg": "Hello World"})
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{} // 0 completed empty

		if err := c.BodyParser(todo); err != nil {
			return err
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body..."})
		}
		//increment

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todo)
	})
	var x int = 5
	var p *int = &x //

	fmt.Println(p)
	fmt.Println(*p)
	/*
		add a Todo
	*/
	//app Patch("/api/todos/:id", func(c *fiber.Ctx) error {
	//	id := c.Params("id")
	//	for i, todo:=	range todos{
	//		if fmt.Sprint(todo.ID) == id{
	//			todos[i].Completed = true
	//			return c.Status(200).JSON(todos[i])
	//		}
	//	}
	//	return c.Status(app.Listen(":4000"))
	//})
	log.Fatal(app.Listen(":4000"))
}
