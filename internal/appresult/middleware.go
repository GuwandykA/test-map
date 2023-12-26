package appresult

import (
	"bd-backend/internal/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type appHandler func(w http.ResponseWriter, r *http.Request) error

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenOriginalMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		fmt.Println("token::::::::::::", token)

		cfg := config.GetConfig()
		claims, err := TokenClaims(token, cfg.JwtKeySupAdmin)
		if err != nil || fmt.Sprint(claims["uuid"]) == "" {
			respondWithError(c, 400, err)
			return
		}

		c.Set("uuid", fmt.Sprint(claims["uuid"]))
		c.Next()
	}
}

func TokenClaims(token, secretKey string) (jwt.MapClaims, error) {
	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		fmt.Println("err", err)
		return nil, ErrMissingParam
	}

	claims, ok := decoded.Claims.(jwt.MapClaims)

	if !ok {
		// TODO tokenin omrini test etmeli
		return nil, ErrInternalServer
	}

	return claims, nil
}
