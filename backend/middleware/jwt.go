package middleware

import (
	"B00k/model"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var Key = []byte("SEVMTE9XT1JELU1PVEhFUkZVQ0tFUg==")

type JWTStruct struct {
	model.UserInfo
	jwt.RegisteredClaims
}

func GenerateToken(user model.UserInfo) (string, error) {
	claims := JWTStruct{
		UserInfo: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "urb00k-system",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(Key)
}

func ParseToken(token_str string) (*JWTStruct, error) {
	token, err := jwt.ParseWithClaims(token_str, &JWTStruct{}, func(token *jwt.Token) (interface{}, error) {
		return Key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTStruct); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func JWTAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHead := c.GetHeader("Authorization")
		if authHead == "" {
			c.JSON(http.StatusBadRequest,
				model.ReturnStruct[int]{
					Status: http.StatusBadRequest,
					Msg:    "Auth failed",
					Data:   1,
				},
			)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHead, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusBadRequest,
				model.ReturnStruct[int]{
					Status: http.StatusBadRequest,
					Msg:    "Authorization header format must be Bearer {token}",
					Data:   1,
				},
			)
			c.Abort()
			return
		}

		claims, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusBadRequest,
				model.ReturnStruct[int]{
					Status: http.StatusBadRequest,
					Msg:    err.Error(),
					Data:   1,
				},
			)
			c.Abort()
			return
		}

		c.Set("user", claims.UserInfo)
		c.Next()
	}
}
