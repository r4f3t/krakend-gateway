package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	jwtMiddleware "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if req.Username != "admin" || req.Password != "password" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": req.Username,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": tokenString})
}

func getOrders(c echo.Context) error {

	authHeader := c.Request().Header.Get("Authorization")
	fmt.Println("Authorization Header:", authHeader) // Hata ayıklama için
	orders := []map[string]interface{}{
		{"id": 1, "item": "Laptop"},
		{"id": 2, "item": "Phone"},
	}
	response := map[string]interface{}{
		"orders": orders,
	}
	return c.JSON(http.StatusOK, response)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/login", login)
	e.GET("/orders", getOrders, jwtMiddleware.WithConfig(jwtMiddleware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	e.Logger.Fatal(e.Start(":8080"))
}
