package utils

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// CustomPayload 自定义载荷继承原有接口并附带自己的字段
type CustomPayload struct {
	UserId uint64
	jwt.RegisteredClaims
}

// GenerateToken 生成Token: uid 用户id, secret 加盐
func GenerateToken(uid uint64, secret string) (string, error) {
	claim := CustomPayload{
		UserId: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "yezi",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))
	return token, err
}

// ParseToken 解析token
func ParseToken(token string, secret string) (*CustomPayload, error) {
	parseToken, err := jwt.ParseWithClaims(token, &CustomPayload{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if parseToken == nil {
		return nil, errors.New("无效的token")
	}
	if claim, ok := parseToken.Claims.(*CustomPayload); ok && parseToken.Valid {
		return claim, nil
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return nil, errors.New("token格式不正确")
	} else if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
		return nil, errors.New("token签名无效")
	} else if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, errors.New("token已过期")
	} else {
		return nil, errors.New("token不可用")
	}
}
