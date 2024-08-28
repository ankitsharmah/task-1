package controller

import (
	
	"clubApi/tasks"
	"net/http"

	// "time"

	"github.com/labstack/echo/v4"
)


func CreateEvent(c echo.Context) error {
    res:=tasks.CreateEvent(c);
	if res!=nil{

		data:= map[string]interface{}{
			"error":res,
		}

		return c.JSON(http.StatusOK,data);
	}
	
	return c.JSON(http.StatusOK,res);

}


func RegisterUserForEvent(c echo.Context) error {
    // Extract the event ID from the URL parameters
    res:=tasks.RegisterUserForEvent(c);
	if res!=nil{

		data:= map[string]interface{}{
			"error":res,
		}

		return c.JSON(http.StatusOK,data);
	}
	
	return c.JSON(http.StatusOK,res);

}

func FindAllEvent(c echo.Context) error {
    res:=tasks.FindAllEvent(c);
	if res!=nil{

		data:= map[string]interface{}{
			"error":res,
		}

		return c.JSON(http.StatusOK,data);
	}
	
	return c.JSON(http.StatusOK,res);

}

