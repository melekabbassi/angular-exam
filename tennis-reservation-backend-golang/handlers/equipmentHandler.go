package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/melekabbassi/tennis-reservation/database"
)

type EquipmentDTO struct {
	Id          int     `json:"id"`
	Type        string  `json:"type"`
	IsAvailable bool    `json:"is_available"`
	Price       float64 `json:"price"`
}

// GET /equipments
func GetEquipments(c *fiber.Ctx) error {
	db := database.OpenDB()

	rows, err := db.Query("SELECT * FROM equipments")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	equipments := make([]EquipmentDTO, 0)

	for rows.Next() {
		equipment := EquipmentDTO{}
		err := rows.Scan(&equipment.Id, &equipment.Type, &equipment.IsAvailable, &equipment.Price)
		if err != nil {
			return err
		}
		equipments = append(equipments, equipment)
	}

	if err = rows.Err(); err != nil {
		return err
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(equipments)
}

// GET /equipments/:id
func GetEquipment(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	equipment := EquipmentDTO{}

	err := db.QueryRow("SELECT * FROM equipments WHERE id = ?", id).Scan(&equipment.Id, &equipment.Type, &equipment.IsAvailable, &equipment.Price)
	if err != nil {
		return c.Status(500).SendString("Equipment not found")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(equipment)
}

// POST /equipments
func CreateEquipment(c *fiber.Ctx) error {
	db := database.OpenDB()

	equipment := EquipmentDTO{}

	if err := c.BodyParser(&equipment); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var lastID int
	err := db.QueryRow("SELECT MAX(id) FROM equipments").Scan(&lastID)
	if err != nil {
		lastID = 0
	}

	_, err = db.Exec("INSERT INTO equipments VALUES (?, ?, ?, ?)", lastID+1, equipment.Type, equipment.IsAvailable, equipment.Price)
	if err != nil {
		return c.Status(500).SendString("Error while creating equipment")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(equipment)
}

// PUT /equipments/:id
func UpdateEquipment(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	equipment := EquipmentDTO{}

	err := c.BodyParser(&equipment)
	if err != nil {
		return c.Status(500).SendString("Error while parsing equipment")
	}

	var equipmentId int
	err = db.QueryRow("SELECT id FROM equipments WHERE id = ?", id).Scan(&equipmentId)
	if err != nil {
		return c.Status(500).SendString("Equipment not found")
	}

	_, err = db.Exec("UPDATE equipments SET type = ?, is_available = ?, price = ? WHERE id = ?", equipment.Type, equipment.IsAvailable, equipment.Price, id)
	if err != nil {
		return c.Status(500).SendString("Error while updating equipment")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(equipment)
}

// DELETE /equipments/:id
func DeleteEquipment(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	var deletedEquipmentId int
	err := db.QueryRow("SELECT id FROM equipments WHERE id = ?", id).Scan(&deletedEquipmentId)
	if err != nil {
		return c.Status(500).SendString("Equipment not found")
	}

	_, err = db.Exec("DELETE FROM equipments WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString("Error while deleting equipment")
	}

	_, err = db.Exec("UPDATE equipments SET id = id - 1 WHERE id > ?", deletedEquipmentId)
	if err != nil {
		return c.Status(500).SendString("Error while updating equipment")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.SendString("Equipment successfully deleted")
}
