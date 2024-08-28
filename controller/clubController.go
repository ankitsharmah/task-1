package controller

import (
	
	"clubApi/tasks"
	"net/http"
	

	"github.com/labstack/echo/v4"
)

func SaveClub(c echo.Context) error{

	res:=tasks.SaveClub(c);
	if res!=nil{

		data:= map[string]interface{}{
			"error":res,
		}

		return c.JSON(http.StatusOK,data);
	}
	
	return c.JSON(http.StatusOK,res);

}

func GetClubById(c echo.Context) error{
	res:=tasks.GetClubById(c);
	if res!=nil{

		data:= map[string]interface{}{
			"error":res,
		}

		return c.JSON(http.StatusOK,data);
	}
	
	return c.JSON(http.StatusOK,res);


}

func GetAllClubs(c echo.Context) error{

	res:=tasks.GetAllClubs(c);
	if res!=nil{

		data:= map[string]interface{}{
			"error":res,
		}

		return c.JSON(http.StatusOK,data);
	}
	
	return c.JSON(http.StatusOK,res);
}



// task
// event api
// findEventByClubId


//table creation 
// clubId ,eventId,userId
