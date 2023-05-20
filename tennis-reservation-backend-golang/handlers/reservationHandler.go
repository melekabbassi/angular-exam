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
	db := database.OpenDB()
	defer database.CloseDB(db)

	reservation := ReservationDTO{}

	if err := c.BodyParser(&reservation); err != nil {
		return c.Status(500).SendString("Error while parsing reservation")
	}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		return c.Status(500).SendString("Error starting transaction")
	}

	// Insert the reservation
	reservationQuery := "INSERT INTO reservations (user_id, court_id, type, date, hour, duration) VALUES (?, ?, ?, ?, ?, ?)"
	reservationStmt, err := tx.Prepare(reservationQuery)
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString("Error preparing reservation statement")
	}
	reservationResult, err := reservationStmt.Exec(reservation.User.Id, reservation.Court.Id, reservation.Type, reservation.Date, reservation.Hour, reservation.Duration)
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString("Error creating reservation")
	}

	// Get the inserted reservation ID
	reservationID, err := reservationResult.LastInsertId()
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString("Error getting reservation ID")
	}

	// Insert the related equipment
	equipmentQuery := "INSERT INTO reservation_equipments (reservation_id, equipment_id) VALUES (?, ?)"
	equipmentStmt, err := tx.Prepare(equipmentQuery)
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString("Error preparing equipment statement")
	}
	for _, equipment := range reservation.Equipment {
		_, err = equipmentStmt.Exec(reservationID, equipment.Id)
		if err != nil {
			tx.Rollback()
			return c.Status(500).SendString("Error creating reservation equipment")
		}
	}

	// Insert the related services
	serviceQuery := "INSERT INTO reservation_services (reservation_id, service_id) VALUES (?, ?)"
	serviceStmt, err := tx.Prepare(serviceQuery)
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString("Error preparing service statement")
	}
	for _, service := range reservation.Services {
		_, err = serviceStmt.Exec(reservationID, service.Id)
		if err != nil {
			tx.Rollback()
			return c.Status(500).SendString("Error creating reservation service")
		}
	}

	// Update court availability
	_, err = tx.Exec("UPDATE courts SET is_available = false WHERE id = ?", reservation.Court.Id)
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString("Error updating court availability")
	}

	// Update equipment availability
	_, err = tx.Exec("UPDATE equipments SET is_available = false WHERE id IN (?)", reservation.GetEquipmentIDs())
	if err != nil {
		tx.Rollback()
		return c.Status(500).SendString("Error updating equipment availability")
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return c.Status(500).SendString("Error committing transaction")
	}

	c.Set("Content-Type", "application/json")
	return c.SendString("Reservation created successfully")
}

// // POST /reservations
// func CreateReservation(c *fiber.Ctx) error {
// 	db := database.OpenDB()

// 	reservation := ReservationDTO{}

// 	if err := c.BodyParser(&reservation); err != nil {
// 		return c.Status(500).SendString("Error while parsing reservation")
// 	}

// 	var lastID int
// 	err := db.QueryRow("SELECT MAX(id) FROM reservations").Scan(&lastID)
// 	if err != nil {
// 		lastID = 0
// 	}

// 	_, err = db.Exec("INSERT INTO reservations (id, user_id, court_id, type, date, hour, duration) VALUES (?, ?, ?, ?, ?, ?, ?)", lastID+1, reservation.User.Id, reservation.Court.Id, reservation.Type, reservation.Date, reservation.Hour, reservation.Duration)
// 	if err != nil {
// 		return c.Status(500).SendString("Error while creating reservation")
// 	}
// 	// if the reservation is created, the related equipments and services are created too
// 	for _, equipmentId := range reservation.Equipment {
// 		_, err = db.Exec("INSERT INTO reservation_equipments (reservation_id, equipment_id) VALUES (?, ?)", lastID+1, equipmentId)
// 		if err != nil {
// 			return c.Status(500).SendString("Error while creating reservation equipment")
// 		}
// 	}

// 	for _, serviceId := range reservation.Services {
// 		_, err = db.Exec("INSERT INTO reservation_services (reservation_id, service_id) VALUES (?, ?)", lastID+1, serviceId)
// 		if err != nil {
// 			return c.Status(500).SendString("Error while creating reservation service")
// 		}
// 	}
// 	// when the reservation is created, the court's availability is set to false and the equipment's availability is set to false
// 	_, err = db.Exec("UPDATE courts SET is_available = false WHERE id = ?", reservation.Court.Id)
// 	if err != nil {
// 		return c.Status(500).SendString("Error while updating court")
// 	}
// 	for _, equipmentId := range reservation.Equipment {
// 		_, err = db.Exec("UPDATE equipments SET is_available = false WHERE id = ?", equipmentId)
// 		if err != nil {
// 			return c.Status(500).SendString("Error while updating equipment")
// 		}
// 	}
// 	database.CloseDB(db)
// 	c.Set("Content-Type", "application/json")

// 	return c.SendString("Reservation created successfully")
// }

// PUT /reservations/:id
func UpdateReservation(c *fiber.Ctx) error {
	db := database.OpenDB()

	reservationId := c.Params("id")

	reservation := ReservationDTO{}

	err := c.BodyParser(&reservation)
	if err != nil {
		return c.Status(500).SendString("Error while parsing reservation")
	}

	var existingReservationId int
	err = db.QueryRow("SELECT id FROM reservations WHERE id = ?", reservationId).Scan(&existingReservationId)
	if err != nil {
		return c.Status(500).SendString("Reservation not found")
	}

	_, err = db.Exec("UPDATE reservations SET user_id = ?, court_id = ?, type = ?, date = ?, hour = ?, duration = ? WHERE id = ?", reservation.User.Id, reservation.Court.Id, reservation.Type, reservation.Date, reservation.Hour, reservation.Duration, reservationId)
	if err != nil {
		return c.Status(500).SendString("Error while updating reservation")
	}

	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.SendString("Reservation updated successfully")
}

// DELETE /reservations/:id
func DeleteReservation(c *fiber.Ctx) error {
	db := database.OpenDB()

	reservationId := c.Params("id")

	var existingReservationId int
	err := db.QueryRow("SELECT id FROM reservations WHERE id = ?", reservationId).Scan(&existingReservationId)
	if err != nil {
		return c.Status(500).SendString("Reservation not found")
	}

	_, err = db.Exec("DELETE FROM reservations WHERE id = ?", reservationId)
	if err != nil {
		return c.Status(500).SendString("Error while deleting reservation")
	}
	// if the reservation is deleted, the related equipments and services are deleted too
	_, err = db.Exec("DELETE FROM reservation_equipments WHERE reservation_id = ?", reservationId)
	if err != nil {
		return c.Status(500).SendString("Error while deleting reservation equipments")
	}
	_, err = db.Exec("DELETE FROM reservation_services WHERE reservation_id = ?", reservationId)
	if err != nil {
		return c.Status(500).SendString("Error while deleting reservation services")
	}
	_, err = db.Exec("UPDATE reservations SET id = id - 1 WHERE id > ?", reservationId)
	if err != nil {
		return c.Status(500).SendString("Error while updating reservations")
	}
	_, err = db.Exec("UPDATE reservation_equipments SET reservation_id = reservation_id - 1 WHERE reservation_id > ?", reservationId)
	if err != nil {
		return c.Status(500).SendString("Error while updating reservation equipments")
	}
	_, err = db.Exec("UPDATE reservation_services SET reservation_id = reservation_id - 1 WHERE reservation_id > ?", reservationId)
	if err != nil {
		return c.Status(500).SendString("Error while updating reservation services")
	}

	// when the reservation is deleted, the court's availability is set to true and the equipment's availability is set to true
	_, err = db.Exec("UPDATE courts SET is_available = true WHERE id = ?", reservationId)
	if err != nil {
		return c.Status(500).SendString("Error while updating court")
	}
	_, err = db.Exec("UPDATE equipments SET is_available = true WHERE id = ?", reservationId)
	if err != nil {
		return c.Status(500).SendString("Error while updating equipment")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.SendString("Reservation deleted")
}

func (r *ReservationDTO) GetEquipmentIDs() []int {
	ids := make([]int, len(r.Equipment))
	for i, equipment := range r.Equipment {
		ids[i] = equipment.Id
	}
	return ids
}
