package auth

import (
	"errors"
	"log"
	database "root/src/database/controller"
	"root/src/database/model"
	"root/src/util"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterUser(c *fiber.Ctx) error {
	type Request struct {
		InitDataRaw string `json:"init_data"`
	}
	user := new(Request)

	if err := c.BodyParser(user); err != nil {
		errResp := &fiber.Map{
			"message": err.Error(),
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return err
	}

	if user.InitDataRaw == "" {
		errResp := &fiber.Map{
			"message": "Init data is required",
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return c.SendStatus(500)
	}

	isValidDataRaw := util.IsUserIitDataValid(user.InitDataRaw)
	if !isValidDataRaw {
		errResp := &fiber.Map{
			"message": "Init data is invalid",
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return c.SendStatus(500)
	}

	userData, err := util.GetDataInInitDataRaw(user.InitDataRaw)

	if err != nil {
		errResp := &fiber.Map{
			"message": err.Error(),
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return c.SendStatus(500)
	}

	userId, err := strconv.Atoi(userData["user_id"])

	if err != nil {
		errResp := &fiber.Map{
			"message": err.Error(),
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return c.SendStatus(500)
	}

	userResponseData := &model.User{
		Id:        uint(userId),
		Name:      userData["username"],
		LastName:  userData["last_name"],
		FirstName: userData["first_name"],
		Hash:      userData["hash"],
	}
	database.DB.Create(&userResponseData)
	return c.Status(200).JSON(userResponseData)
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
