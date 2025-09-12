package main

import (
    "context"
    "fmt"
    "log"
	"strconv"

    "github.com/gofiber/fiber/v2"
    db "gamecraft-backend/prisma/generated/prisma-client" // alias db
)

func main() {
    client := db.NewClient()
    if err := client.Prisma.Connect(); err != nil {
        log.Fatal(err)
    }
    defer client.Prisma.Disconnect()

    app := fiber.New()

    app.Get("/ping", func(c *fiber.Ctx) error {
        return c.SendString("pong")
    })

    // Create user
    app.Post("/user", func(c *fiber.Ctx) error {
        type Request struct {
            Name string `json:"name"`
        }
        req := new(Request)
        if err := c.BodyParser(req); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "invalid request"})
        }

        user, err := client.User.CreateOne(
            db.User.Name.Set(req.Name),
        ).Exec(context.Background())
        if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": err.Error()})
        }

        return c.JSON(user)
    })

    // Get user by ID
    app.Get("/user/:id", func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
		}

		user, err := client.User.FindUnique(
			db.User.ID.Equals(id),
		).Exec(context.Background())
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "user not found"})
		}
		return c.JSON(user)
	})


    fmt.Println("Server running on http://localhost:3000")
    log.Fatal(app.Listen(":3000"))
}
