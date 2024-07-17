package improvement

import (
	database "root/src/database/controller"
	"root/src/database/model"

	"github.com/gofiber/fiber/v2"
)

func AddUserImprovement(c *fiber.Ctx) error {
	var req struct {
		UserName      string `json:"user_name"`
		ImprovementID int    `json:"improvement_id"`
		Value         int    `json:"value"`
	}

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
