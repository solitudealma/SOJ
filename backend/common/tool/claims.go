package tool

import (
	"context"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UserId   int64
	Username string
}

func GetClaims(signingKey string, ctx context.Context) (*CustomClaims, error) {

	token, ok := ctx.Value("token").(string)
	if !ok {
		return nil, errors.New("请求携带的token类型不正确")
	}

	j := NewJWT(signingKey)
	claims, err := j.ParseToken(token)
	if err != nil {
		logx.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}

	return claims, err
}

// GetUserId 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserId(signingKey string, ctx context.Context) int64 {

	claims, ok := ctx.Value("claims").(*CustomClaims)
	if ok {
		cl, err := GetClaims(signingKey, ctx)
		if err != nil {
			return 0
		}

		return cl.UserId
	}

	return claims.UserId
}
