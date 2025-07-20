package mytools

import (
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

var jwtSecret = []byte("thisIsMysteriousData")

type Claims struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name"`
	Authority int    `json:"authority"` // 权限
	jwt.RegisteredClaims
}

// CreateToken 生成Token (有效期为一天)
func CreateToken(id int, userName string, authority int) (string, error) {
	nowTime := time.Now()
	expirationTime := nowTime.Add(12 * 3600 * time.Second)

	claims := Claims{
		ID:        id,
		UserName:  userName,
		Authority: authority,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(nowTime),        // 签发时间
			ExpiresAt: jwt.NewNumericDate(expirationTime), // 过期时间
			NotBefore: jwt.NewNumericDate(nowTime),        // 生效时间(立即生效)
			Issuer:    "myapp",                            // 签发者标识
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析并验证Token
func ParseToken(tokenString string) (*Claims, error) {
	tokenS := strings.Split(tokenString, " ")[1]
	token, err := jwt.ParseWithClaims(tokenS, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}

/*--------------------------------------------------------------------------------------------------------------------*/

type EmailClaims struct {
	UserID        int    `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType int    `json:"operation_type"`
	jwt.RegisteredClaims
}

// CreateEmailToken 生成email-Token (有效期为一小时)
func CreateEmailToken(userId, operation int, email, password string) (string, error) {
	nowTime := time.Now()
	expirationTime := nowTime.Add(3600 * time.Second)

	claims := EmailClaims{
		UserID:        userId,
		Email:         email,
		Password:      password,
		OperationType: operation,

		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(nowTime),        // 签发时间
			ExpiresAt: jwt.NewNumericDate(expirationTime), // 过期时间
			NotBefore: jwt.NewNumericDate(nowTime),        // 生效时间(立即生效)
			Issuer:    "myapp",                            // 签发者标识
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseEmailToken 解析并验证邮件Token
func ParseEmailToken(tokenString string) (*EmailClaims, error) {
	tokenS := strings.Split(tokenString, " ")[1]
	token, err := jwt.ParseWithClaims(tokenS, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if token != nil {
		if claims, ok := token.Claims.(*EmailClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
