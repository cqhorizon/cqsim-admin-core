package storage

import (
	"log"

	"cqsim-admin-core/sdk/config"
	"cqsim-admin-core/sdk/pkg/captcha"
)

// Setup 配置storage组件
func Setup() {
	//4. 设置缓存
	cacheAdapter, err := config.CacheConfig.Setup()
	if err != nil {
		log.Fatalf("cache setup error, %s\n", err.Error())
	}
	//5. 设置验证码store
	captcha.SetStore(captcha.NewCacheStore(cacheAdapter, 600))
}
