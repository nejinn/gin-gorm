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

// 获取传入token只有user数据
func GetJwtUser(c *gin.Context) (res model.User) {
	user, _ := c.Get("user")
	userInfo := user.(model.User)
	return userInfo
}

// token验证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.ErrorMsgWithResponseMsg(c, response.TokenError)
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == global.NLY_CONFIG.Jwt.TokenPrefix) {
			response.ErrorWithCustomMsg(c, response.TokenError.Code, "token错误")
			c.Abort()
			return
		}
		// 这个步骤一定要有
		j := &JWT{
			SignKey: []byte(global.NLY_CONFIG.Jwt.SignKey),
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		claims, err := j.ParseToken(parts[1])
		if err != nil {
			response.ErrorWithCustomMsg(c, response.TokenError.Code, err.Error())
			c.Abort()
			return
		}
		var user model.User
		err = global.NLY_DB.Where("username=?", claims.Username).First(&user).Error
		if err != nil {
			response.ErrorWithCustomMsg(c, response.TokenError.Code, "token错误")
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
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
	}
	return nil, errors.New("认证信息不对")
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
			Issuer:    global.NLY_CONFIG.Jwt.TokenPrefix,
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
