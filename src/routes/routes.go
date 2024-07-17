package routes

import (
	admin "root/src/routes/Admin"
	amplifiers "root/src/routes/Amplifiers"
	"root/src/routes/auth"

	"github.com/gofiber/fiber/v2"
)

func AllRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("auth/register", auth.RegisterUser)
	api.Post("auth/login", auth.LoginUser)
	api.Get("auth/user/:userId", auth.GetUserById)

	api.Post("amplifiers/saveClicks", amplifiers.SaveClicks)

	adminDh := api.Group("/admin")

	adminDh.Post("improvements/create", admin.CreateImprovement)
	adminDh.Post("addUserImprovement", admin.AddUserImprovement)
	adminDh.Get("getUserImprovements", admin.GetUserImprovements) //*
	adminDh.Post("deleteUserImprovement", admin.DeleteUserImprovement)
	adminDh.Post("deleteImprovement", admin.DeleteImprovement)
}
