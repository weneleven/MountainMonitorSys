package model

import "github.com/dgrijalva/jwt-go"

type MyClaims struct { //定义要求
	Username string `json:"username"`
	jwt.StandardClaims
}
