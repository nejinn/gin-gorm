package config

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtConf struct {
	SignKey         string `yaml:"sign-key"`
	TokenExpireTime int    `yaml:"token-expire-time"`
	ExpireTimeUnit  string `yaml:"expire-time-unit"`
	TokenPrefix     string `yaml:"token-prefix"`
}

type JwtClaims struct {
	Username string
	jwt.StandardClaims
}
