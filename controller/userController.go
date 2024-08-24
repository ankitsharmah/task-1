package controller

import (
	"clubApi/config"
	"clubApi/models"
	"fmt"
	"net/http"
	"strconv"

	// "strings"

	"github.com/labstack/echo/v4"
)

// SaveUser handles the saving of a user to the database.
func SaveUser(c echo.Context) error {
	u := new(models.User) // Correct usage

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	db := config.DB()

	if err := db.Create(u).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to save user")
	}

	response := map[string]interface{}{
		"data": u,
	}

	return c.JSON(http.StatusOK, response)
}


func GetAllUser (c echo.Context) error{

	var users[] *models.User

	db:=config.DB()

	if res := db.Find(&users); res.Error != nil {
		data := map[string]string{
			"message":"faild in sending users" ,
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": users,
	} 

	return c.JSON(http.StatusOK, response)

}

func GetUserById(c echo.Context) error {
	idStr := c.Param("id")
	fmt.Printf("Received ID: %s\n", idStr) // Debug print

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid user ID")
	}

	db := config.DB()

	user := new(models.User)

	if err := db.First(user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, user)
}

