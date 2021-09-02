package request

import (
	"github.com/golang-jwt/jwt/v4"
)

// CustomClaims structure
type CustomClaims struct {
	Username    string
	BufferTime  int64
	jwt.StandardClaims
}
