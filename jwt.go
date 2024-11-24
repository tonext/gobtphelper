package gobtphelper

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	UserId int64 `json:"userId"`
	jwt.StandardClaims
}

func GetToken(userId int64) string {
	// 创建一个自定义的声明
	claims := CustomClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(), // 令牌有效期为 72 小时
			Issuer:    "btp.ccszhd.com",
		},
	}

	secretKey := GetConfig("jwt_secret_key")

	// 使用 HS256 算法生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Println("Error generating token:", err)
		return ""
	}
	return tokenString
}

func GetUserIdByToken(tokenString string) (int64, error) {
	secretKey := GetConfig("jwt_secret_key")
	// 解析 JWT
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 检查签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, fmt.Errorf("error parsing token: %v", err)
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		// 检查是否过期
		if claims.ExpiresAt < time.Now().Unix() {
			return 0, fmt.Errorf("token has expired")
		}
		return claims.UserId, nil
	}
	return 0, fmt.Errorf("invalid token")
}
