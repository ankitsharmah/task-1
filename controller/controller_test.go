package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"

	// "strconv"
	"testing"

	"clubApi/config"
	"clubApi/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSaveUser(t *testing.T) {
	e := echo.New()
	config.DatabaseInit()

	user := models.User{Name: "John Doe", Phone: "1234567890"}
	body, _ := json.Marshal(user)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/users/add", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Set the content type
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, SaveUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]interface{}
		if err := json.Unmarshal(rec.Body.Bytes(), &response); assert.NoError(t, err) {
			assert.Equal(t, "John Doe", response["data"].(map[string]interface{})["name"])
		}
	}
}


func TestGetAllUser(t *testing.T) {
	e := echo.New()
	config.DatabaseInit()

	// Add test user
	user := models.User{Name: "Jane Doe", Phone: "0987654321"}
	db := config.DB()
	db.Create(&user)

	// Create a new GET request
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set the handler and call it
	if assert.NoError(t, GetAllUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]interface{}
		if err := json.Unmarshal(rec.Body.Bytes(), &response); assert.NoError(t, err) {
			users := response["data"].([]interface{})
			assert.True(t, len(users) > 0)
		}
	}
}

func TestGetUserById(t *testing.T) {
	e := echo.New()
	config.DatabaseInit()

	// Create a test user
	user := models.User{Name: "Jack Doe", Phone: "1122334455"}
	db := config.DB()
	result := db.Create(&user)
	if result.Error != nil {
		t.Fatalf("Failed to create user: %v", result.Error)
	}

	// Use the dynamically retrieved ID
	userID := user.ID

	// Print the userID for debugging
	fmt.Printf("Testing GetUserById with userID: %d\n", userID)

	// Create a new GET request with the correct user ID
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/"+strconv.Itoa(userID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Manually set the route parameter in the context
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(userID))

	// Set the handler and call it
	if assert.NoError(t, GetUserById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response models.User
		if err := json.Unmarshal(rec.Body.Bytes(), &response); assert.NoError(t, err) {
			// Check if the response matches the user created
			assert.Equal(t, user.Name, response.Name)
			assert.Equal(t, user.Phone, response.Phone)
		}
	}
}




func TestGetClubById(t *testing.T) {
	e := echo.New()
	config.DatabaseInit()

	// Create a test user
	club := models.Club{ClubName: "Club b", Address: "tirupati"}
	db := config.DB()
	result := db.Create(&club)
	if result.Error != nil {
		t.Fatalf("Failed to create club: %v", result.Error)
	}

	// Use the dynamically retrieved ID
	clubID := club.ID

	// Print the userID for debugging
	fmt.Printf("Testing GetUserById with club: %d\n", clubID)

	// Create a new GET request with the correct user ID
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users/"+strconv.Itoa(clubID), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Manually set the route parameter in the context
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(clubID))

	// Set the handler and call it
	if assert.NoError(t, GetClubById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response models.Club
		if err := json.Unmarshal(rec.Body.Bytes(), &response); assert.NoError(t, err) {
			// Check if the response matches the user created
			assert.Equal(t, club.ClubName, response.ClubName)
			assert.Equal(t, club.Address, response.Address)
		}
	}
}

func TestSaveClub(t *testing.T) {
	e := echo.New()
	config.DatabaseInit()

	// Test data
	club := models.Club{ClubName: "Club A", Address: "123 Main St"}
	body, _ := json.Marshal(club)

	// Create a new POST request
	req := httptest.NewRequest(http.MethodPost, "/api/v1/club/add", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON) // Set the content type

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set the handler and call it
	if assert.NoError(t, SaveClub(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]interface{}
		if err := json.Unmarshal(rec.Body.Bytes(), &response); assert.NoError(t, err) {
			assert.Equal(t, "club saved successfully", response["message"])
		}
	}
}

func TestGetAllClubs(t *testing.T) {
	e := echo.New()
	config.DatabaseInit()

	// Add test club
	club := models.Club{ClubName: "Club B", Address: "456 Elm St"}
	db := config.DB()
	db.Create(&club)

	// Create a new GET request
	req := httptest.NewRequest(http.MethodGet, "/api/v1/club/all", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set the handler and call it
	if assert.NoError(t, GetAllClubs(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]interface{}
		if err := json.Unmarshal(rec.Body.Bytes(), &response); assert.NoError(t, err) {
			clubs := response["clubs"].([]interface{})
			assert.True(t, len(clubs) > 0)
		}
	}
}
