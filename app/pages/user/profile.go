package user

import (
	"GoHtmxGL/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Routes(api fiber.Router) {
	userGroup := api.Group("/users")
	{
		userGroup.Get("/", getUsers)
		userGroup.Get("/:id", getUser)
		userGroup.Post("/", createUser)
		userGroup.Put("/:id", updateUser)
		userGroup.Delete("/:id", deleteUser)
	}
}

//todo: convert to htmx templates

// getUsers handles the GET request to fetch all user
func getUsers(ctx *fiber.Ctx) error {
	// Logic to fetch and return all users
	fmt.Println("Getting users")
	var users []models.User
	//initializers.GetDB().Find(&users)
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": users})
}

// getUser handles the GET request to fetch a specific user
func getUser(ctx *fiber.Ctx) error {
	// Logic to fetch and return a specific user
	id := ctx.Params("id")
	var user models.User
	if result := models.DB.First(&user, id); result.Error != nil {
		// Handle the error, such as returning an error response
		return ctx.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})
}

// createUser handles the POST request to create a new user
func createUser(ctx *fiber.Ctx) error {
	var user models.User
	// Get user from context
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}
	// Create the user in the models
	if err := models.DB.Create(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}
	// Return the created user in the response
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})
}

// updateUser handles the PUT request to update an user
func updateUser(ctx *fiber.Ctx) error {
	// Get user ID from the request parameters
	userID := ctx.Params("id")

	// Parse JSON request body into User struct
	var updatedUser models.User
	if err := ctx.BodyParser(&updatedUser); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})

	}
	// Fetch the existing user from the models
	var existingUser models.User
	if err := models.DB.First(&existingUser, userID).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	// Update the user fields
	existingUser.Username = updatedUser.Username
	existingUser.Role = updatedUser.Role
	//existingUser.Password = updatedUser.Password

	// Save the updated user to the models
	//if err := initializers.GetDB().Save(&existingUser).Error; err != nil {
	//	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	//}

	// Return the updated user in the response
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": updatedUser})
}

// deleteUser handles the DELETE request to delete an user
func deleteUser(ctx *fiber.Ctx) error {
	// Get user ID from the request parameters
	userID := ctx.Params("id")

	// Fetch the existing user from the models
	var existingUser models.User
	if err := models.DB.First(&existingUser, userID).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	//Delete the user from the models
	if err := models.DB.Delete(&existingUser).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	// Return a success response
	return ctx.JSON(fiber.Map{"message": "User deleted successfully"})
}
