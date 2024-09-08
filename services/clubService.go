package tasks

import (
	"clubApi/config"
	"clubApi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SaveClub(c echo.Context) error {

	club := new(*models.Club)
	db := config.DB()
	if err := c.Bind(club); err != nil {
		return err
	}

	if err := db.Create(club).Error; err != nil {

		data := map[string]interface{}{
			"error": err,
		}

		return c.JSON(http.StatusOK, data)
	}
	res := map[string]string{
		"message": "club saved successfully",
	}
	return c.JSON(http.StatusOK, res)

}


func GetClubById(c echo.Context) error{
	club :=new(*models.Club)
	db:=config.DB()

	id,err:=strconv.Atoi(c.Param("id"))
	
	if err!=nil {
		return err
	}
	if err := db.Preload("Events.Users").First(&club, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Club not found")
	}
	

	return c.JSON(http.StatusOK,club)


}

func GetAllClubs(c echo.Context) error{

	db:=config.DB()

	var clubs[] *models.Club



	if err:=db.Find(&clubs).Error; err!=nil{
		data:=map[string]string{
			"message":"error in finding all clubs",
		}

		return c.JSON(http.StatusOK,data)
	}

	res:=map[string]interface{}{
		"clubs":clubs,
	}

	return c.JSON(http.StatusOK,res)
}