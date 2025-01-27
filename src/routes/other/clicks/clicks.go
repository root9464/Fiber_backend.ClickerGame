package clicks

import (
	"log"
	database "root/src/database/controller"
	"root/src/database/model"

	"github.com/gofiber/fiber/v2"
)

func SaveClicks(c *fiber.Ctx) error {
	type Request struct {
		UserId int `json:"user_id"`
		Clicks int `json:"clicks"`
	}

	request := new(Request)

	if err := c.BodyParser(request); err != nil {
		errResp := &fiber.Map{
			"message": err.Error(),
		}
		if err := c.Status(500).JSON(errResp); err != nil {
			log.Fatal(err)
		}
		return err
	}

	database.DB.Model(&model.ProgressClicker{}).Where("user_id = ?", request.UserId).Update("clicks", request.Clicks)
	return c.Status(200).JSON(fiber.Map{"message": "Success"})
}
