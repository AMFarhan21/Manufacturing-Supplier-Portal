package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/AMFarhan21/fres"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func forbiddenResponse(c echo.Context) error {
	return c.JSON(http.StatusForbidden, fres.Response.StatusForbidden(http.StatusForbidden))
}

func JWTMiddleware(jwtSecret string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			signature := strings.Split(c.Request().Header.Get("Authorization"), " ")
			if len(signature) < 2 {
				return forbiddenResponse(c)
			}

			if signature[0] != "Bearer" {
				return forbiddenResponse(c)
			}

			claims := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(signature[1], claims, func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
				}

				return []byte(jwtSecret), nil
			})

			if err != nil {
				return forbiddenResponse(c)
			}

			method, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok || method != jwt.SigningMethodHS256 {
				return forbiddenResponse(c)
			}

			expAt, err := claims.GetExpirationTime()
			if err != nil {
				return forbiddenResponse(c)
			}

			if time.Now().After(expAt.Time) {
				return forbiddenResponse(c)
			}

			userId := claims["id"].(string)
			role := claims["role"].(string)

			c.Set("id", userId)
			c.Set("role", role)
			return next(c)
		}
	}
}

func ACLMiddleware(rolesMap map[string]bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role, _ := c.Get("role").(string)
			if role == "admin" {
				return next(c)
			}

			if rolesMap[role] {
				return next(c)
			}

			return forbiddenResponse(c)
		}
	}
}
