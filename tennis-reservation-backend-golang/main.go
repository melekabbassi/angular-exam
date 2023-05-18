package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/melekabbassi/tennis-reservation/database"
	"github.com/melekabbassi/tennis-reservation/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	loadEnv()

	app := generateApp()

	db := database.OpenDB()
	defer database.CloseDB(db)

	app.Listen(":8083")
}

func loadEnv() error {
	goENV := os.Getenv("GO_ENV")
	if goENV == "" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}

func generateApp() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	userGroup := app.Group("/users")
	userGroup.Get("/", handlers.GetUsers)
	userGroup.Get("/:id", handlers.GetUser)
	userGroup.Post("/", handlers.CreateUser)
	userGroup.Put("/:id", handlers.UpdateUser)
	userGroup.Delete("/:id", handlers.DeleteUser)

	return app
}
