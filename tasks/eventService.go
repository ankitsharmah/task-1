package tasks

import (
	"clubApi/config"
	"clubApi/models"
	"fmt"
	"net/http"
	"strconv"
	// "time"

	"github.com/labstack/echo/v4"
)

func CreateEvent(c echo.Context) error {
    db := config.DB()
    event := new(models.Event)

    if err := c.Bind(event); err != nil {
        return c.JSON(http.StatusBadRequest, err)
    }

    clubID, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid club ID")
    }

    // Get current date and remove time
    // event.Date = time.Now().Truncate(24 * time.Hour)

    // event.Date = parsedDate
    event.ClubID = uint(clubID)

    if err := db.Create(&event).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, err)
    }

    return c.JSON(http.StatusCreated, event)
}

func RegisterUserForEvent(c echo.Context) error {
    // Extract the event ID from the URL parameters
	db:=config.DB()
    eventIDStr := c.Param("id")
    eventID, err := strconv.Atoi(eventIDStr)
    uid := c.Param("userid")
    uu, err := strconv.Atoi(uid)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid event ID format")
    }

    // Bind the request body to a User object
    var user models.User
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, err.Error())
    }

    fmt.Printf("Event ID: %d\n", eventID)

    // Find the event by ID
    var event models.Event
    if err := db.First(&event, eventID).Error; err != nil {
        fmt.Printf("Error finding event: %v\n", err)
        return c.JSON(http.StatusNotFound, "Event not found")
    }


    if err := db.Find(&user,uu).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to find or create user")
    }

    if err := db.Model(&event).Association("Users").Append(&user); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to register user for event")
    }

    return c.JSON(http.StatusOK, user)
}

func FindAllEvent(c echo.Context) error {
    db := config.DB()
    var events []models.Event

    // Use Preload to load associated Users with each Event
    if err := db.Preload("Users").Find(&events).Error; err != nil {
        return c.JSON(http.StatusNotFound, err)
    }

    return c.JSON(http.StatusOK, events)
}