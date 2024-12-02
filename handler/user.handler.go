package handler

import (
	"go-fiber/database"
	"go-fiber/models/entity"
	"go-fiber/models/request"
	"log"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error{
	var users []entity.User
	result := database.DB.Debug().Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)
}


var validate = validator.New()

func UserHandlerCreate(ctx *fiber.Ctx) error{
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}
	
	
		// Validasi input menggunakan validator
	if errValidate := validate.Struct(user); errValidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errValidate.Error(),
		})
	}
	

	newUser := entity.User{
		Name: user.Name,
		Email: user.Email,
		Address: user.Address,
		Phone: user.Phone,
	}

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"massage": "failed to store data",
		})
	}
		// Mengembalikan respons sukses
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Success",
		"data":    newUser,
	})

}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	// Ambil parameter ID dari URL
	userId := ctx.Params("id")

	var user entity.User
	// Query untuk mencari user berdasarkan ID
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		// Jika tidak ditemukan, kembalikan status 404
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// Jika ditemukan, kembalikan data user
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
		"data":    user,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	//menangkap data reques 
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"Massage": "bad request",
		})
	}

	// Ambil parameter ID dari URL
	userId := ctx.Params("id")
	var user entity.User
	// Query untuk mencari user berdasarkan ID
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		// Jika tidak ditemukan, kembalikan status 404
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	//Update Data
	if userRequest.Name != ""{
		user.Name = userRequest.Name
	}
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone


	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"massage": "internal server erorr",
		})
	}
	return ctx.JSON(fiber.Map{
		"massage": "Success Updated",
		"data" : user,
	})


}

func UserHandlerUpdateEmail(ctx *fiber.Ctx) error {
	//menangkap data reques 
	userRequest := new(request.UserEmailRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"Massage": "bad request",
		})
	}

	// Ambil parameter ID dari URL
	userId := ctx.Params("id")
	//membuat variabel user mengambil dari models
	var user entity.User
	// Query untuk mencari user berdasarkan ID
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		// Jika tidak ditemukan, kembalikan status 404
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}


	// chek apakah ada email yg sama
	var userEmail entity.User
	errEmail := database.DB.First(&userEmail, "email = ?", userRequest.Email)

	if errEmail == nil {
		return ctx.Status(402).JSON( fiber.Map{
			"message":"email already used",
		})
	}

	// update Email
	user.Email = userRequest.Email
	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"massage": "internal server erorr",
		})
	}
	return ctx.JSON(fiber.Map{
		"massage": "Success Updated Email",
		"data" : user,
	})
	
}

func UserHandlerDeleted(ctx *fiber.Ctx) error {
	// Ambil parameter ID dari URL
	userId := ctx.Params("id") 
	//membuat variabel user mengambil dari models
	var user entity.User
	// Query untuk mencari user berdasarkan ID
	err := database.DB.Delete(&user, "id = ?", userId).Error
	// Jika ada error saat menghapus
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to delete user",
		})
	}

	// Berhasil dihapus
	return ctx.Status(200).JSON(fiber.Map{
		"message": "user deleted successfully",
	})
}