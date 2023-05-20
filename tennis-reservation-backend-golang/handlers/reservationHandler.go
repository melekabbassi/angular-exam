package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/melekabbassi/tennis-reservation/database"
	"github.com/melekabbassi/tennis-reservation/models"
)

type ReservationDTO struct {
	Id        int                `json:"id"`
	User      models.User        `json:"user"`
	Court     models.Court       `json:"court"`
	Type      string             `json:"type"`
	Date      string             `json:"date"`
	Hour      int                `json:"hour"`
	Duration  int                `json:"duration"`
	Equipment []models.Equipment `json:"equipment"`
	Services  []models.Service   `json:"services"`
}

// POST /reservations
// func CreateReservation(c *fiber.Ctx) error {
// 	db := database.OpenDB()

// 	reservation := ReservationDTO{}

// 	if err := c.BodyParser(&reservation); err != nil {
// 		return c.Status(500).SendString("Error while parsing reservation")
// 	}

// 	userID := reservation.User.Id

// 	var exisingUserID int
// 	err := db.QueryRow("SELECT id FROM users WHERE id = ?", userID).Scan(&exisingUserID)
// 	if err != nil {
// 		return c.Status(500).SendString("User not found")
// 	}

// 	for _, equipmentID := range reservation.Equipment {
// 		var existingEquipmentID int
// 		err := db.QueryRow("SELECT id FROM equipments WHERE id = ?", equipmentID).Scan(&existingEquipmentID)
// 		if err != nil {
// 			return c.Status(500).SendString("Equipment not found")
// 		}
// 	}

// 	for _, serviceID := range reservation.Services {
// 		var existingServiceID int
// 		err := db.QueryRow("SELECT id FROM services WHERE id = ?", serviceID).Scan(&existingServiceID)
// 		if err != nil {
// 			return c.Status(500).SendString("Service not found")
// 		}
// 	}

// 	_, err = db.Exec("INSERT INTO reservations (user_id, court_id, type, date, hour, duration) VALUES (?, ?, ?, ?, ?, ?)", userID, reservation.Court.Id, reservation.Type, reservation.Date, reservation.Hour, reservation.Duration)
// 	if err != nil {
// 		return c.Status(500).SendString("Error while creating reservation")
// 	}

// 	var reservationID int
// 	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&reservationID)
// 	if err != nil {
// 		return c.Status(500).SendString("Error while getting reservation id")
// 	}

// 	for _, equipmentID := range reservation.Equipment {
// 		_, err = db.Exec("INSERT INTO reservation_equipments (reservation_id, equipment_id) VALUES (?, ?)", reservationID, equipmentID)
// 		if err != nil {
// 			return c.Status(500).SendString("Error while creating reservation equipment")
// 		}
// 	}

// 	for _, serviceID := range reservation.Services {
// 		_, err = db.Exec("INSERT INTO reservation_services (reservation_id, service_id) VALUES (?, ?)", reservationID, serviceID)
// 		if err != nil {
// 			return c.Status(500).SendString("Error while creating reservation service")
// 		}
// 	}
// 	database.CloseDB(db)
// 	c.Set("Content-Type", "application/json")

// 	return c.JSON(reservation)
// }

// GET /reservations
func GetReservations(c *fiber.Ctx) error {
	db := database.OpenDB()

	userId := c.Params("id")

	rows, err := db.Query("SELECT * FROM reservations WHERE user_id = ?", userId)
	if err != nil {
		return c.Status(500).SendString("Error while getting reservations")
	}

	var reservations []ReservationDTO

	for rows.Next() {
		var reservation ReservationDTO
		if err := rows.Scan(&reservation.Id, &reservation.User, &reservation.Court, &reservation.Type, &reservation.Date, &reservation.Hour, &reservation.Duration); err != nil {
			return c.Status(500).SendString("Error while getting reservation")
		}
		reservations = append(reservations, reservation)
	}

	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(reservations)
}

// GET /reservations/:id
func GetReservation(c *fiber.Ctx) error {
	db := database.OpenDB()

	reservationId := c.Params("id")

	var reservation ReservationDTO
	if err := db.QueryRow("SELECT * FROM reservations WHERE id = ?", reservationId).Scan(&reservation.Id, &reservation.User, &reservation.Court, &reservation.Type, &reservation.Date, &reservation.Hour, &reservation.Duration); err != nil {
		return c.Status(500).SendString("Reservation not found")
	}

	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(reservation)
}

// POST /reservations
func CreateReservation(c *fiber.Ctx) error {
	if err := validateUser(c); err != nil {
		return err
	}

	if err := validateEquipment(c); err != nil {
		return err
	}

	if err := validateService(c); err != nil {
		return err
	}

	return saveReservation(c)
}

// DELETE /reservations/:id
func DeleteReservation(c *fiber.Ctx) error {
	db := database.OpenDB()

	reservationId := c.Params("id")

	_, err := db.Exec("DELETE FROM reservations WHERE id = ?", reservationId)
	if err != nil {
		return c.Status(500).SendString("Error while deleting reservation")
	}

	database.CloseDB(db)
	return nil
}

func validateUser(c *fiber.Ctx) error {
	db := database.OpenDB()

	userID := c.Params("id")

	var exisingUserID int
	err := db.QueryRow("SELECT id FROM users WHERE id = ?", userID).Scan(&exisingUserID)
	if err != nil {
		return c.Status(500).SendString("User not found")
	}

	database.CloseDB(db)
	return nil
}

func validateEquipment(c *fiber.Ctx) error {
	db := database.OpenDB()

	equipmentId := c.Params("id")

	var existingEquipmentId int
	err := db.QueryRow("SELECT id FROM equipments WHERE id = ?", equipmentId).Scan(&existingEquipmentId)
	if err != nil {
		return c.Status(500).SendString("Equipment not found")
	}

	database.CloseDB(db)
	return nil
}

func validateService(c *fiber.Ctx) error {
	db := database.OpenDB()

	serviceId := c.Params("id")

	var existingServiceId int
	err := db.QueryRow("SELECT id FROM services WHERE id = ?", serviceId).Scan(&existingServiceId)
	if err != nil {
		return c.Status(500).SendString("Service not found")
	}

	database.CloseDB(db)
	return nil
}

func saveReservation(c *fiber.Ctx) error {
	db := database.OpenDB()

	reservation := ReservationDTO{}

	if err := c.BodyParser(&reservation); err != nil {
		return c.Status(500).SendString("Error while parsing reservation")
	}

	userID := reservation.User.Id

	_, err := db.Exec("INSERT INTO reservations (user_id, court_id, type, date, hour, duration) VALUES (?, ?, ?, ?, ?, ?)", userID, reservation.Court.Id, reservation.Type, reservation.Date, reservation.Hour, reservation.Duration)
	if err != nil {
		return c.Status(500).SendString("Error while creating reservation")
	}

	var reservationId int
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&reservationId)
	if err != nil {
		return c.Status(500).SendString("Error while getting reservation id")
	}

	for _, equipmentId := range reservation.Equipment {
		_, err = db.Exec("INSERT INTO reservation_equipments (reservation_id, equipment_id) VALUES (?, ?)", reservationId, equipmentId)
		if err != nil {
			return c.Status(500).SendString("Error while creating reservation equipment")
		}
	}

	for _, serviceId := range reservation.Services {
		_, err = db.Exec("INSERT INTO reservation_services (reservation_id, service_id) VALUES (?, ?)", reservationId, serviceId)
		if err != nil {
			return c.Status(500).SendString("Error while creating reservation service")
		}
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(reservation)
}
