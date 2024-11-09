package services

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"learn.com/config"
	"learn.com/models"
	"net/http"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword), err
}

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	fmt.Println(user.Password)
	hashedPassword, err := HashPassword(user.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	user.Password = hashedPassword

	if err := config.DB.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

func LoginUser(c echo.Context) error {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var user models.User
	if err := config.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		fmt.Println("Password comparison error:", err)
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Incorrect password"})
	}

	token, err := GenerateJWT(user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
func GetUsers(c echo.Context) error {
	var users []models.User
	config.DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}

func GetUserData(c echo.Context) error {
	userIDValue := c.Get("user_id")

	userID, ok := userIDValue.(uint)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid or missing user ID"})
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}
