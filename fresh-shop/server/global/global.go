package global

import (
	"sync"

	"fresh-shop/server/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"fresh-shop/server/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB
	DbList             map[string]*gorm.DB
	Redis              *redis.Client
	Config             config.Server
	Viper              *viper.Viper
	Log                *zap.Logger
	Timer              timer.Timer = timer.NewTimerTask()
	ConcurrencyControl             = &singleflight.Group{}
	SugarLog           *zap.SugaredLogger
	BlackCache         local_cache.Cache
	lock               sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DbList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DbList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
