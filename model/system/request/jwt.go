package request

import (
	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

// CustomClaims structure
type CustomClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	BufferTime  int64
	jwt.StandardClaims
}
