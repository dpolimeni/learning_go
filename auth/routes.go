package auth

import (
	"context"
	"fmt"

	"github.com/dpolimeni/fiber_app/common"
	"github.com/dpolimeni/fiber_app/ent/user"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var DbClient = common.GetDB()

// Register godoc
// @Summary Register a person on DB.
// @Tags Authentication
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/auth/register
// @Router /api/v1/auth/register [post]
// @Param NewUser body NewUser true "New User to register"
func Register(c *fiber.Ctx) error {
	newUser := new(NewUser)
	if err := c.BodyParser(newUser); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("Error hashing password")
	}
	newUser.Password = string(hashedPassword)
	user, err := DbClient.User.Create().
		SetUsername(newUser.Username).
		SetPassword(newUser.Password).
		SetEmail(newUser.Email).
		Save(context.Background())
	if err != nil {
		error_msg := fmt.Sprintf("Error creating user: %s", err)
		return c.Status(fiber.StatusBadRequest).JSON(error_msg)
	}
	user_response := UserResponse{
		Username: user.Username,
		Email:    user.Email,
	}
	return c.Status(fiber.StatusCreated).JSON(user_response)
}

// Login godoc
// @Summary Login a person on DB.
// @Tags Authentication
// @Accept application/json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @BasePath /api/v1/auth
// @Router /api/v1/auth/login [post]
// @Param UserLogin body NewUser true "Login form"
func Login(c *fiber.Ctx) error {
	userLogin := new(UserLogin)
	if err := c.BodyParser(userLogin); err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	user, err := DbClient.User.Query().Where(user.UsernameEQ(userLogin.Username)).First(context.Background())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON("User not found")
	}
	fmt.Println(user)
	// err = bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(user.Password))
	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).JSON("Invalid Password")
	// }
	return c.Status(fiber.StatusAccepted).JSON("Logged in")
}

func SetUpAuthRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Post("/auth/register", Register)
	v1.Post("/auth/login", Login)
}
