package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json::"title"`
	DONE  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Print("Hello World")

	app := fiber.New()

	todos := []Todo{}

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		// catch error, test error, if find return error
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		// first ID will be 1 since by default no TODO is in system
		todo.ID = len(todos) + 1

		todos = append(todos, *todo) //appending a pointer

		return c.JSON(todos) // then return the latest todo

	})

	app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("Id")

		if err != nil {
			return c.Status(401).SendString("Invalid id")
		}

		for i, t := range todos {
			if t.ID == id {
				todos[i].DONE = true
				break
			}
		}

		return c.JSON(todos)
	})

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	// if app.Listen fails, it is going to trigger log.Fatal
	log.Fatal(app.Listen(":4000"))

}
