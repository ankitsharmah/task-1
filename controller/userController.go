package controller

import (
	
	"clubApi/tasks"
	"net/http"

	// "strings"

	"github.com/labstack/echo/v4"
)

// SaveUser handles the saving of a user to the database.
func SaveUser(c echo.Context) error {

	response:=tasks.SaveUser(c);


	if response!= nil {
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, response)
}


func GetAllUser (c echo.Context) error{

	res:=tasks.GetAllUser(c);

	if res != nil {
		data := map[string]string{
			"message":"faild in sending users" ,
		}

		return c.JSON(http.StatusOK, data)
	}

	response := map[string]interface{}{
		"data": res,
	} 

	return c.JSON(http.StatusOK, response)

}

func GetUserById(c echo.Context) error {

	res:=tasks.GetUserById(c);

	if res!= nil {
		return echo.NewHTTPError(http.StatusNotFound,res)
	}

	return c.JSON(http.StatusOK, res)
}

