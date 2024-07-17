package routes

import (
	admin "root/src/routes/Admin"

	"root/src/routes/auth"
	"root/src/routes/other/clicks"
	improvement "root/src/routes/other/improvement"

	"github.com/gofiber/fiber/v2"
)

func AllRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("auth/register", auth.RegisterUser)
	api.Post("auth/login", auth.LoginUser)
	api.Get("auth/user/:userId", auth.GetUserById)

	api.Post("improvement/addUserImprovement", improvement.AddUserImprovement)
	api.Get("improvement/getImprovements", improvement.GetImprovements)
	api.Post("saveClicks", clicks.SaveClicks)
	adminDh := api.Group("/admin")

	adminDh.Post("improvements/create", admin.CreateImprovement)
	adminDh.Post("deleteUserImprovement", admin.DeleteUserImprovement)
	adminDh.Post("deleteImprovement", admin.DeleteImprovement)
}
