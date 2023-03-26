package middleware

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/solitudealma/SOJ/backend/app/comment/cmd/api/internal/config"
	"github.com/solitudealma/SOJ/backend/common/result"
	"github.com/solitudealma/SOJ/backend/common/tool"
	"github.com/solitudealma/SOJ/backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
	"golang.org/x/sync/singleflight"
)

type AuthMiddleware struct {
	Config             config.Config
	ConcurrencyControl *singleflight.Group
	RefisClient        *redis.Redis
}

func NewAuthMiddleware(c config.Config, concurrencyControl *singleflight.Group, redis *redis.Redis) *AuthMiddleware {
	return &AuthMiddleware{
		Config:             c,
		ConcurrencyControl: concurrencyControl,
		RefisClient:        redis,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 我们这里jwt鉴权取头部信息 token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := r.Header.Get("token")
		if token == "" {
			resp := result.ResponseErrorBean{
				Code: xerr.REUQEST_PARAM_ERROR,
				Msg:  "未登录或非法访问",
			}

			httpx.WriteJson(w, http.StatusOK, resp)

			return
		}

		if _, err := m.RefisClient.GetCtx(r.Context(), token); err != nil {
			resp := result.ResponseErrorBean{
				Code: xerr.REUQEST_PARAM_ERROR,
				Msg:  "您的账号异地登录或令牌失效",
			}

			httpx.WriteJson(w, http.StatusOK, resp)

			return
		}

		j := tool.NewJWT(m.Config.Auth.AccessSecret)
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == tool.TokenExpired {
				resp := result.ResponseErrorBean{
					Code: xerr.REUQEST_PARAM_ERROR,
					Msg:  "授权已过期",
				}

				httpx.WriteJson(w, http.StatusOK, resp)

				return
			}

			resp := result.ResponseErrorBean{
				Code: xerr.REUQEST_PARAM_ERROR,
				Msg:  err.Error(),
			}

			httpx.WriteJson(w, http.StatusOK, resp)

			return
		}

		if int64(claims.ExpiresAt.Sub(time.Now())) < claims.BufferTime {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(m.Config.Auth.AccessExpire)))
			newToken, _ := j.CreateTokenByOldToken(m.ConcurrencyControl, token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			r.Header.Set("new-token", newToken)
			r.Header.Set("new-expires-at", strconv.FormatInt(int64(newClaims.ExpiresAt.Second()), 10))

			_, err := m.RefisClient.GetCtx(r.Context(), newClaims.Username)
			if err != nil {
				logx.WithContext(r.Context()).Error("get redis jwt failed", err)
			} else { // 当之前的取成功时才进行拉黑操作
				_, _ = m.RefisClient.DelCtx(r.Context(), newClaims.Username)
			}
			// 无论如何都要记录当前的活跃状态
			_ = m.RefisClient.SetCtx(r.Context(), newClaims.Username, newToken)

		}

		r = r.WithContext(context.WithValue(r.Context(), "token", token))
		r = r.WithContext(context.WithValue(r.Context(), "claims", claims))
		next(w, r)
	}
}
