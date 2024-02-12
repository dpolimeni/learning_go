package main

import (
	"context"
	"fmt"

	"github.com/dpolimeni/fiber_app/auth"
	"github.com/dpolimeni/fiber_app/common"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Create a superuser
	new_user := auth.NewUser{
		Username: "admin",
		Password: "admin",
		Email:    "admin@gmail.com",
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(new_user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	new_user.Password = string(hashedPassword)
	DbClient := common.GetDB()
	// Register the superuser
	user, err := DbClient.User.Create().SetUsername(new_user.Username).SetPassword(new_user.Password).SetEmail(new_user.Email).SetIsAdmin(true).Save(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("USER CREATED:")
	fmt.Println(user)
}
