/*
 * @Author: fzf404
 * @Date: 2022-02-25 11:28:13
 * @LastEditTime: 2022-02-27 14:56:30
 * @Description: JWT 鉴权
 */

package utils

import (
	"gover/global"
	"gover/model/user"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	SignKey []byte // 密钥
}

type JWTClaims struct {
	UserID uint // 用户 ID
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		SignKey: []byte(global.CONFIG.JWT.Key),
	}
}

/**
 * @description: 生成 Token
 * @param {*user.User} u
 */
func (j *JWT) GenerateToken(u *user.User) (string, error) {
	claims := JWTClaims{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                      // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.Expires, // 签名过期时间
			Issuer:    global.CONFIG.JWT.Issuer,                      // 签名发行人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SignKey)
}

/**
 * @description: 验证 token
 * @param {string} tokenString
 */
func (j *JWT) ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SignKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
