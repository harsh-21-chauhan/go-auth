package helpers


import (
	"errors"
  "time"
  "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-super-secret-key")


func GenerateToken(userID uint, role string) (string, error) {
  claims := jwt.MapClaims{
    "user_id": userID,
    "role":    role,
    "exp":     time.Now().Add(24 * time.Hour).Unix(),
  }
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(jwtSecret)
}

func ValidateToken(tokenStr string) (jwt.MapClaims, error) {
  token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
    return jwtSecret, nil
  })



  if err != nil || !token.Valid {
    return nil, err
  }



  claims, ok := token.Claims.(jwt.MapClaims)
  if !ok {
    return nil, errors.New("invalid claims")
  }
  return claims, nil
}