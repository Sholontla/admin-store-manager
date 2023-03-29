package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

const secretKey = "9ikNKMPJQ8dNIsDKQ9zaGbjH4Zp5IU0NA=="

type ClaimsWithScope struct {
	jwt.RegisteredClaims
	Scope       string
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

func IsAuthenticated(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")
	claims := ClaimsWithScope{}
	token, err := jwt.ParseWithClaims(cookie, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthenticated"})
	}

	payLoad := token.Claims.(*ClaimsWithScope)

	isSuperAdmin := strings.Contains(ctx.Path(), "/service/super/admin")
	if (payLoad.Scope == "admin" && isSuperAdmin) || (payLoad.Scope == "super_admin" && !isSuperAdmin) {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Unauthorized"})
	}

	isAdmin := strings.Contains(ctx.Path(), "/service/admin")
	if (payLoad.Scope == "super_admin" && isAdmin) || (payLoad.Scope == "admin" && !isAdmin) {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "Unauthorized"})
	}

	return ctx.Next()
}

func GetUserLogin(c *fiber.Ctx) (string, error) {
	cookie := c.Cookies("jwt")

	claims := ClaimsWithScope{}

	token, err := jwt.ParseWithClaims(cookie, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}

	payLoad := token.Claims.(*ClaimsWithScope)

	return payLoad.Subject, nil
}

func GenerateJWT(adminEmail string, scope string, role []string, permissions []string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)
	payLoad := ClaimsWithScope{}
	payLoad.Subject = adminEmail
	payLoad.ExpiresAt = jwt.NewNumericDate(expireTime)
	payLoad.Scope = scope
	payLoad.Roles = role
	payLoad.Permissions = permissions
	return jwt.NewWithClaims(jwt.SigningMethodHS256, payLoad).SignedString([]byte(secretKey))

}
