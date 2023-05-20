package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/melekabbassi/tennis-reservation/database"
)

type CourtDTO struct {
	Id          int  `json:"id"`
	IsAvailable bool `json:"is_available"`
}

// GET /courts
func GetCourts(c *fiber.Ctx) error {
	db := database.OpenDB()

	rows, err := db.Query("SELECT * FROM courts")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	courts := make([]CourtDTO, 0)

	for rows.Next() {
		court := CourtDTO{}
		err := rows.Scan(&court.Id, &court.IsAvailable)
		if err != nil {
			return err
		}
		courts = append(courts, court)
	}

	if err = rows.Err(); err != nil {
		return err
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(courts)
}

// GET /courts/:id
func GetCourt(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	court := CourtDTO{}

	err := db.QueryRow("SELECT * FROM courts WHERE id = ?", id).Scan(&court.Id, &court.IsAvailable)
	if err != nil {
		return c.Status(500).SendString("Court not found")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(court)
}

// POST /courts
func CreateCourt(c *fiber.Ctx) error {
	db := database.OpenDB()

	court := CourtDTO{}

	if err := c.BodyParser(&court); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	_, err := db.Exec("INSERT INTO courts (is_available) VALUES (?)", court.IsAvailable)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(court)
}

// PUT /courts/:id
func UpdateCourt(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	court := CourtDTO{}

	err := c.BodyParser(&court)
	if err != nil {
		return c.Status(500).SendString("Error while parsing court")
	}

	var courtId int
	err = db.QueryRow("SELECT id FROM courts WHERE id = ?", id).Scan(&courtId)
	if err != nil {
		return c.Status(500).SendString("Court not found")
	}

	_, err = db.Exec("UPDATE courts SET is_available = ? WHERE id = ?", court.IsAvailable, id)
	if err != nil {
		return c.Status(500).SendString("Error while updating court")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(court)
}

// DELETE /courts/:id
func DeleteCourt(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	var deletedCourtId int
	err := db.QueryRow("SELECT id FROM courts WHERE id = ?", id).Scan(&deletedCourtId)
	if err != nil {
		return c.Status(500).SendString("Court not found")
	}

	_, err = db.Exec("DELETE FROM courts WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString("Error while deleting court")
	}

	_, err = db.Exec("UPDATE courts SET id = id - 1 WHERE id > ?", deletedCourtId)
	if err != nil {
		return c.Status(500).SendString("Error while updating courts")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.SendString("Court successfully deleted")
}
