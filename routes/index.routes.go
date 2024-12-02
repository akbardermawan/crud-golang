package routes

import (
	"go-fiber/config"
	"go-fiber/handler"

	"github.com/gofiber/fiber/v2"
)

//membuat middleware
func Middleware(ctx *fiber.Ctx) error{
	// contoh penggunaan middleware sebagai autentification
	token := ctx.Get("x-token")
	if token != "secret"{
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Massage": "Unauthorized",
		})
	}
	
	return ctx.Next()
}

func RouteInit(r *fiber.App){
	// membuat route untuk mengakses static asset
	r.Static("/public",config.ProjectRootPath + "/public/asset")

	// membuat route 
	r.Get("/user", Middleware, handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Put("/user/email/:id", handler.UserHandlerUpdateEmail)
	r.Delete("/user/:id", handler.UserHandlerDeleted)
}