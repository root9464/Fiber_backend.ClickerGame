package admin

import (
	"log"
	database "root/src/database/controller"
	"root/src/database/model"

	"github.com/gofiber/fiber/v2"
)

func CreateImprovement(c *fiber.Ctx) error {
	request := new(model.Improvement)

	if err := c.BodyParser(request); err != nil {
		errResp := &fiber.Map{
			"message": err.Error(),
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return err
	}

	database.DB.Create(&request)
	return c.Status(200).JSON(fiber.Map{"message": "Success"})
}

func AddUserImprovement(c *fiber.Ctx) error {
	var req struct {
		UserName      string `json:"user_name"`
		ImprovementID int    `json:"improvement_id"`
		Value         int    `json:"value"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid request body"})
	}

	var user model.User
	err := database.DB.Where("name = ?", req.UserName).First(&user).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "user not found"})
	}

	var improvement model.Improvement
	err = database.DB.Where("id = ?", req.ImprovementID).First(&improvement).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "improvement not found"})
	}

	// Проверяем, существует ли уже улучшение для этого пользователя
	var existingUserImprovement model.UserImprovement
	err = database.DB.Where("user_id = ? AND improvement_id = ?", user.Id, improvement.Id).First(&existingUserImprovement).Error
	if err == nil { // Улучшение уже существует, обновляем его
		existingUserImprovement.Value = req.Value
		err = database.DB.Save(&existingUserImprovement).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error updating user improvement"})
		}
	} else { // Создаем новое улучшение для пользователя
		userImprovement := &model.UserImprovement{
			UserId:        user.Id,
			UserName:      user.Name,
			ImprovementId: improvement.Id,
			Improvement:   improvement,
			Value:         req.Value,
		}
		err = database.DB.Create(userImprovement).Error
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error creating user improvement"})
		}
	}

	return c.Status(200).JSON(fiber.Map{"message": "success"})
}

func GetUserImprovements(c *fiber.Ctx) error {
	userImprovements := make([]model.UserImprovement, 0)
	err := database.DB.Find(&userImprovements).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "error getting user improvements"})
	}
	return c.Status(200).JSON(userImprovements)
}

func DeleteUserImprovement(c *fiber.Ctx) error {
	var request struct {
		UserName      string `json:"user_name"`
		ImprovementID int    `json:"improvement_id"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid request body"})
	}
	user := new(model.User)
	err := database.DB.Where("name = ?", request.UserName).First(&user).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "user not found"})
	}
	improvement := new(model.Improvement)

	err = database.DB.Where("id = ?", request.ImprovementID).First(&improvement).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "improvement not found"})
	}
	err = database.DB.Where("user_id = ? AND improvement_id = ?", user.Id, improvement.Id).Delete(&model.UserImprovement{}).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "error deleting user improvement"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "success"})
}

func DeleteImprovement(c *fiber.Ctx) error {
	var request struct {
		ImprovementID int `json:"improvement_id"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid request body"})
	}
	improvement := new(model.Improvement)

	err := database.DB.Where("id = ?", request.ImprovementID).First(&improvement).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "improvement not found"})
	}
	err = database.DB.Delete(&improvement).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "error deleting improvement"})
	}
	return c.Status(200).JSON(fiber.Map{"message": "success"})
}
