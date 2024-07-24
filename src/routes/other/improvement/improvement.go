package improvement

import (
	"log"
	database "root/src/database/controller"
	"root/src/database/model"

	"github.com/gofiber/fiber/v2"
)

func AddUserImprovement(c *fiber.Ctx) error {
	type Request struct {
		UserName      string `json:"user_name"`
		ImprovementID int    `json:"improvement_id"`
		Value         int    `json:"value"`
	}

	req := new(Request)

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid request body"})
	}

	tx := database.DB.Begin()
	defer tx.Rollback()

	var user model.User
	if err := tx.Where("name = ?", req.UserName).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "user not found"})
	}

	var improvement model.Improvement
	if err := tx.Where("id = ?", req.ImprovementID).First(&improvement).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "improvement not found"})
	}

	var existingUserImprovement model.UserImprovement
	if err := tx.Where("user_id = ? AND improvement_id = ?", user.Id, improvement.Id).First(&existingUserImprovement).Error; err == nil {
		existingUserImprovement.Value = req.Value
		if err := tx.Save(&existingUserImprovement).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error updating user improvement"})
		}
	} else {
		userImprovement := &model.UserImprovement{
			UserId:        user.Id,
			UserName:      user.Name,
			ImprovementId: improvement.Id,
			Improvement:   improvement,
			Value:         req.Value,
		}
		if err := tx.Create(userImprovement).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "error creating user improvement"})
		}
	}

	tx.Commit()
	return c.Status(200).JSON(fiber.Map{"message": "success"})
}

func GetImprovements(c *fiber.Ctx) error {
	improvements := new([]model.Improvement)
	err := database.DB.Find(&improvements).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "error getting user improvements"})
	}
	return c.Status(200).JSON(improvements)
}

func GetUserImprovements(c *fiber.Ctx) error {
	userId := c.Params("userId")

	if userId == "" {
		errResp := &fiber.Map{
			"message": "Improvement id is required",
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return c.SendStatus(500)
	}

	userImprovements := new([]model.UserImprovement)
	err := database.DB.Where("user_id = ?", userId).Find(&userImprovements).Error

	if err != nil {
		errResp := &fiber.Map{
			"message": "Error getting user improvements",
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return c.SendStatus(500)

	}

	return c.Status(200).JSON(userImprovements)

}
