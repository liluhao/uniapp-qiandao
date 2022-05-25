package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

var (
	ErrMissingHeader = errors.New("未携带token")
)

// Context token令牌;把用户id封装在Context里
type Context struct {
	ID string
}

func ParseRequest(c *gin.Context) (*Context, error) {
	// 获取头部的token
	headerToken := c.Request.Header.Get("Authorization")

	jwtSecret := viper.GetString("jwt_secret")
	// 如果访问接口没有携带token直接返回异常信息
	if len(headerToken) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string
	// 解析标头，获取令牌部分
	fmt.Sscanf(headerToken, "Bearer %s", &t)
	return Parse(t, jwtSecret)
}

// Parse 判断token是否合法
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}
	// 验证，并返回一个令牌。
	// keyFunc 将接收已解析的令牌并应返回用于验证的密钥。如果一切都是 kosher，err 将是 nil
	token, err := jwt.Parse(tokenString, secretFunc(secret))
	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = claims["id"].(string)
		return ctx, nil
	} else {
		return ctx, err
	}
}

// 验证密码格式
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

// Sign 生成token
func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}
	// 定义token令牌内容
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": c.ID,
		// token生效时间
		"nbf": time.Now().Unix(),
		// token签发时间
		"iat": time.Now().Unix(),
		// token过期时间
		"exp": time.Now().Add(time.Hour + 9).Unix(),
	})
	// 对令牌进行签名
	tokenString, err = token.SignedString([]byte(secret))
	return
}
