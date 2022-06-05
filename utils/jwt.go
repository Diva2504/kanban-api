package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secret = "abcd"

// Generate Json Web Token
func GenerateToken(id uint, email string) (string, error) {
  claims := jwt.MapClaims{
    "id" : id,
    "email": email,
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
  signedToken, err := token.SignedString([]byte(secret))
  if err != nil {
    fmt.Errorf("Something went wrong: %s", err.Error())
    return "", err
  }
  return signedToken, nil
}

func VerifyToken(c *gin.Context) {

}
