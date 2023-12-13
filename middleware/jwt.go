package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(c *fiber.Ctx) error {
	token, ok := c.GetReqHeaders()["Authorization"]
	if !ok {
		return fmt.Errorf("Unauthorized")
	}

	claims, err := validateToken(token[0])
	if err != nil {
		return err
	}

	expires := claims["expires"].(float64)
	if time.Now().Unix() > int64(expires) {
		return fmt.Errorf("Token expired")
	}

	return c.Next()
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Invalid signin method", token.Header["alg"])
			return nil, fmt.Errorf("Unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("Failed to parse JWT token: ", err)
		return nil, fmt.Errorf("Unhautorized")
	}

	if !token.Valid {
		fmt.Println("Invalid token")
		return nil, fmt.Errorf("Unhautorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Unhautorized")
	}
	return claims, nil
}
