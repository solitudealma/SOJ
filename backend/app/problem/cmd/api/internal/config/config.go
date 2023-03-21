package config

import (
	"github.com/solitudealma/SOJ/backend/common/stores/gormx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
		BufferTime   int64
	}
	MySQL gormx.Config
	Redis struct {
		Host string
		Pass string
		Type string
	}
	Cache        cache.CacheConf
	JudgeRpcConf zrpc.RpcClientConf
}
