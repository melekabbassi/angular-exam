package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/melekabbassi/tennis-reservation/database"
)

type ServiceDTO struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// GET /services
func GetServices(c *fiber.Ctx) error {
	db := database.OpenDB()

	rows, err := db.Query("SELECT * FROM services")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	services := make([]ServiceDTO, 0)

	for rows.Next() {
		service := ServiceDTO{}
		err := rows.Scan(&service.Id, &service.Name, &service.Description, &service.Price)
		if err != nil {
			return err
		}
		services = append(services, service)
	}

	if err = rows.Err(); err != nil {
		return err
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(services)
}

// GET /services/:id
func GetService(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	service := ServiceDTO{}

	err := db.QueryRow("SELECT * FROM services WHERE id = ?", id).Scan(&service.Id, &service.Name, &service.Description, &service.Price)
	if err != nil {
		return c.Status(500).SendString("Service not found")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(service)
}

// POST /services
func CreateService(c *fiber.Ctx) error {
	db := database.OpenDB()

	service := ServiceDTO{}

	if err := c.BodyParser(&service); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	_, err := db.Exec("INSERT INTO services (name, description, price) VALUES (?, ?, ?)", service.Name, service.Description, service.Price)
	if err != nil {
		return c.Status(500).SendString("Error creating new service")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(service)
}

// PUT /services/:id
func UpdateService(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	service := ServiceDTO{}

	err := c.BodyParser(&service)
	if err != nil {
		return c.Status(500).SendString("Error while parsing service")
	}

	var serviceId int
	err = db.QueryRow("SELECT id FROM services WHERE id = ?", id).Scan(&serviceId)
	if err != nil {
		return c.Status(500).SendString("Service not found")
	}

	_, err = db.Exec("UPDATE services SET name = ?, description = ?, price = ? WHERE id = ?", service.Name, service.Description, service.Price, id)
	if err != nil {
		return c.Status(500).SendString("Error updating service")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(service)
}

// DELETE /services/:id
func DeleteService(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	var deletedServiceId int
	err := db.QueryRow("SELECT id FROM services WHERE id = ?", id).Scan(&deletedServiceId)
	if err != nil {
		return c.Status(500).SendString("Service not found")
	}

	_, err = db.Exec("DELETE FROM services WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString("Error deleting service")
	}

	_, err = db.Exec("UPDATE services SET id = id - 1 WHERE id > ?", id)
	if err != nil {
		return c.Status(500).SendString("Error updating services")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.SendString("Service successfully deleted")
}
