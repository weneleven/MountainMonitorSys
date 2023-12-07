package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mountain/utils"
	"mountain/utils/errmessage"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct { //定义要求
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func Settoken(username string) (string, int) {
	expertime := time.Now().Add(time.Hour * 10) //设置十个小时的有效期
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expertime.Unix(), //有效期至 返回Unix时间戳
			Issuer:    "Mountain",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey) //返回签好名的token
	if err != nil {
		return "", errmessage.ERROR
	}
	return token, errmessage.SUCCESS
}

// 验证token
func Checktoken(token string) (*MyClaims, int) {
	// 先解析拿到jwttoken的指针对象，
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	// 然后调用Vaild方法判断是否成功效验并获取jwt中的Clains类型转换断言为自定义Clamins
	if key, _ := setToken.Claims.(*MyClaims); setToken.Valid {
		return key, errmessage.SUCCESS
	} else {
		return nil, errmessage.ERROR
	}
}

// jwt中间件  固定写法
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		var code int
		if tokenHerder == "" {
			code = errmessage.ERROR_TOKEN_NOT_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmessage.Geterrmessage(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHerder, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmessage.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmessage.Geterrmessage(code),
			})
			c.Abort()
			return
		}
		key, tCode := Checktoken(checkToken[1])
		if tCode != errmessage.SUCCESS {
			code = errmessage.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmessage.Geterrmessage(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmessage.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmessage.Geterrmessage(code),
			})
			c.Abort()
			return
		}

		c.Set("username", key.Username)
		c.Next()
	}
}
