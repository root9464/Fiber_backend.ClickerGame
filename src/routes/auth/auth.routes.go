package auth

import (
	"errors"
	"log"
	database "root/src/database/controller"
	"root/src/database/model"
	"root/src/util"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterUser(c *fiber.Ctx) error {
	user := new(model.User)

	if err := c.BodyParser(user); err != nil {
		errResp := &fiber.Map{
			"message": err.Error(),
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return err
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		errResp := &fiber.Map{
			"message": "All fields are required",
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return c.SendStatus(500)
	}

	if !util.IsEmailValid(user.Email) {
		errResp := &fiber.Map{
			"message": "Invalid email",
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return c.SendStatus(500)
	}

	database.DB.Create(&user)

	return c.Status(200).JSON(user)
}

func LoginUser(c *fiber.Ctx) error {
	type Request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	requestLoginUser := new(Request)

	if err := c.BodyParser(requestLoginUser); err != nil {
		errResp := &fiber.Map{
			"message": err.Error(),
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return err
	}

	if requestLoginUser.Email == "" || requestLoginUser.Password == "" {
		errResp := &fiber.Map{
			"message": "All fields are required",
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return c.SendStatus(500)
	}

	user := new(model.User)
	if err := database.DB.Where("email = ? AND password = ?", requestLoginUser.Email, requestLoginUser.Password).First(&user).Error; err != nil {
		errResp := &fiber.Map{
			"message": "Invalid email or password",
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(user)
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("userId")
	if id == "" {
		errResp := &fiber.Map{
			"message": "Id is required",
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return c.SendStatus(500)
	}

	var user model.User

	response := database.DB.First(&user, "id = ?", id)

	if response.Error != nil {
		if errors.Is(response.Error, gorm.ErrRecordNotFound) {
			errResp := &fiber.Map{
				"message": "User not found",
			}
			if err := c.Status(500).JSON(errResp); err != nil {
				log.Fatal(err)
			}
			return c.SendStatus(500)
		}
	}

	if err := c.Status(200).JSON(user); err != nil {
		log.Fatal(err)
	}

	return c.Status(200).JSON(user)
}
