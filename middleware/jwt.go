package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"note/config"
	"note/global"
	"note/model"
	"note/response"
	"strings"
	"time"
)

type JWT struct {
	SignKey []byte
}

// 构造 生成token方法
func (j *JWT) GenToken(claims config.JwtClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SignKey)
}

// 构造 解析token 方法
func (j *JWT) ParseToken(tokenString string) (*config.JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &config.JwtClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SignKey, nil
	})
	if err != nil {
		if detail, ok := err.(*jwt.ValidationError); ok {
			if detail.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token已过期")
			} else {
				return nil, errors.New("认证信息不对")
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*config.JwtClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("认证信息不对")
	} else {
		return nil, errors.New("认证信息不对")
	}
}

// 构造 过期时间转换 方法
func (j *JWT) GetExpiresNum() int {
	var expiresNum int
	if strings.Compare(global.NLY_CONFIG.Jwt.ExpireTimeUnit, "day") == 0 {
		expiresNum = 60 * 60 * 24 * global.NLY_CONFIG.Jwt.TokenExpireTime
		return expiresNum
	} else if strings.Compare(global.NLY_CONFIG.Jwt.ExpireTimeUnit, "hour") == 0 {
		expiresNum = 60 * 60 * global.NLY_CONFIG.Jwt.TokenExpireTime
		return expiresNum
	} else if strings.Compare(global.NLY_CONFIG.Jwt.ExpireTimeUnit, "minute") == 0 {
		expiresNum = 60 * global.NLY_CONFIG.Jwt.TokenExpireTime
		return expiresNum
	} else {
		expiresNum = global.NLY_CONFIG.Jwt.TokenExpireTime
		return expiresNum
	}
}

// 分发token
func DispenseToken(c *gin.Context, user *model.User) {
	j := &JWT{
		SignKey: []byte(global.NLY_CONFIG.Jwt.SignKey),
	}
	expiresNum := j.GetExpiresNum()
	claims := config.JwtClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 5000,
			ExpiresAt: time.Now().Unix() + int64(expiresNum),
			Issuer:    "nejinn",
		},
	}
	token, err := j.GenToken(claims)
	if err != nil {
		response.ErrorWithCustomMsg(c, response.TokenError.Code, err.Error())
		return
	}
	type Date struct {
		Token string `json:"token"`
	}
	date := Date{
		Token: token,
	}
	response.OkData(c, response.Ok, date)
}
