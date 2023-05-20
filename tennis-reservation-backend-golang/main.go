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

	courtGroup := app.Group("/courts")
	courtGroup.Get("/", handlers.GetCourts)
	courtGroup.Get("/:id", handlers.GetCourt)
	courtGroup.Post("/", handlers.CreateCourt)
	courtGroup.Put("/:id", handlers.UpdateCourt)
	courtGroup.Delete("/:id", handlers.DeleteCourt)

	equipmentGroup := app.Group("/equipments")
	equipmentGroup.Get("/", handlers.GetEquipments)
	equipmentGroup.Get("/:id", handlers.GetEquipment)
	equipmentGroup.Post("/", handlers.CreateEquipment)
	equipmentGroup.Put("/:id", handlers.UpdateEquipment)
	equipmentGroup.Delete("/:id", handlers.DeleteEquipment)

	serviceGroup := app.Group("/services")
	serviceGroup.Get("/", handlers.GetServices)
	serviceGroup.Get("/:id", handlers.GetService)
	serviceGroup.Post("/", handlers.CreateService)
	serviceGroup.Put("/:id", handlers.UpdateService)
	serviceGroup.Delete("/:id", handlers.DeleteService)

	reservationGroup := app.Group("/reservations")
	reservationGroup.Get("/", handlers.GetReservations)
	reservationGroup.Get("/:id", handlers.GetReservation)
	reservationGroup.Post("/", handlers.CreateReservation)
	reservationGroup.Put("/:id", handlers.UpdateReservation)
	reservationGroup.Delete("/:id", handlers.DeleteReservation)

	return app
}
