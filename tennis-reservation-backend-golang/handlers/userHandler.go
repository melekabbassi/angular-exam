package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/melekabbassi/tennis-reservation/database"
	"golang.org/x/crypto/bcrypt"
)

type UserDTO struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsActive  bool   `json:"isActive"`
	Role      string `json:"role"`
}

// GET /users
func GetUsers(c *fiber.Ctx) error {
	db := database.OpenDB()

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	users := make([]UserDTO, 0)

	for rows.Next() {
		user := UserDTO{}
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.IsActive, &user.Role)
		if err != nil {
			return err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return err
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(users)
}

// GET /users/:id
func GetUser(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	user := UserDTO{}

	err := db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.IsActive, &user.Role)
	if err != nil {
		return c.Status(500).SendString("User not found")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.JSON(user)
}

// POST /users
func CreateUser(c *fiber.Ctx) error {
	db := database.OpenDB()

	user := UserDTO{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).SendString("Error while parsing user")
	}

	var email string
	err = db.QueryRow("SELECT email FROM users WHERE email = ?", user.Email).Scan(&email)
	if err == nil {
		return c.Status(500).SendString("User already exists")
	}

	// Encode password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).SendString("Error while hashing password")
	}

	var lastID int
	err = db.QueryRow("SELECT MAX(id) FROM users").Scan(&lastID)
	if err != nil {
		return c.Status(500).SendString("Error while getting last id")
	}

	_, err = db.Exec("INSERT INTO users VALUES (?, ?, ?, ?, ?, ?, ?)", lastID+1, user.FirstName, user.LastName, user.Email, string(hashedPassword), user.IsActive, user.Role)
	if err != nil {
		return c.Status(500).SendString("Error while creating user")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.SendString("User created successfully")
}

// PUT /users/:id
func UpdateUser(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	user := UserDTO{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).SendString("Error while parsing user")
	}

	var userID int
	err = db.QueryRow("SELECT id FROM users WHERE id = ?", id).Scan(&userID)
	if err != nil {
		return c.Status(500).SendString("User not found")
	}

	// if the user didn't change his password, we don't need to hash it again
	if user.Password == "" {
		_, err = db.Exec("UPDATE users SET first_name = ?, last_name = ?, email = ?, is_active = ?, role = ? WHERE id = ?", user.FirstName, user.LastName, user.Email, user.IsActive, user.Role, id)
		if err != nil {
			return c.Status(500).SendString("Error while updating user")
			// return c.Status(500).SendString(err.Error())
		}
	} else {
		// Encode password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).SendString("Error while hashing password")
		}

		_, err = db.Exec("UPDATE users SET first_name = ?, last_name = ?, email = ?, password = ?, is_active = ?, role = ? WHERE id = ?", user.FirstName, user.LastName, user.Email, string(hashedPassword), user.IsActive, user.Role, id)
		if err != nil {
			return c.Status(500).SendString("Error while updating user")
		}
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.SendString("User updated successfully")
}

// // PUT /users/:id
// func UpdateUser(c *fiber.Ctx) error {
// 	db := database.OpenDB()

// 	id := c.Params("id")

// 	user := UserDTO{}

// 	err := c.BodyParser(&user)
// 	if err != nil {
// 		return c.Status(500).SendString("Error while parsing user")
// 	}

// 	var userID int
// 	err = db.QueryRow("SELECT id FROM users WHERE id = ?", id).Scan(&userID)
// 	if err != nil {
// 		return c.Status(500).SendString("User not found")
// 	}

// 	// Encode password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return c.Status(500).SendString("Error while hashing password")
// 	}

// 	fmt.Println(user.FirstName, user.LastName, user.Email, string(hashedPassword), user.IsActive, user.Role, id)
// 	_, err = db.Exec("UPDATE users SET firstName = ?, lastName = ?, email = ?, password = ?, isActive = ?, role = ? WHERE id = ?", user.FirstName, user.LastName, user.Email, string(hashedPassword), user.IsActive, user.Role, id)
// 	if err != nil {
// 		return c.Status(500).SendString("Error while updating user")
// 	}
// 	database.CloseDB(db)
// 	c.Set("Content-Type", "application/json")

// 	return c.SendString("User updated successfully")
// }

// DELETE /users/:id
func DeleteUser(c *fiber.Ctx) error {
	db := database.OpenDB()

	id := c.Params("id")

	var deletedID int
	err := db.QueryRow("SELECT id FROM users WHERE id = ?", id).Scan(&deletedID)
	if err != nil {
		return c.Status(500).SendString("User not found")
	}

	_, err = db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return c.Status(500).SendString("Error while deleting user")
	}

	_, err = db.Exec("UPDATE users SET id = id - 1 WHERE id > ?", deletedID)
	if err != nil {
		return c.Status(500).SendString("Error while updating user ids")
	}
	database.CloseDB(db)
	c.Set("Content-Type", "application/json")

	return c.SendString("User deleted successfully")
}
