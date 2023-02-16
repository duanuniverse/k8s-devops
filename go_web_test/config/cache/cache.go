package cache

import (
	"fmt"
	"github.com/allegro/bigcache"
	log "github.com/sirupsen/logrus"
	"go_web_test/utils"
	"math"
	"time"
)

var BigCache *Cache

// Cache 缓存
type Cache struct {
	BigCache *bigcache.BigCache // 本地缓存
}

// Get 根据key从缓存中获取对象
func (c Cache) Get(key string) (value interface{}, err error) {
	valueBytes, err := c.BigCache.Get(key)
	if err != nil {
		return nil, err
	}
	value = utils.Deserialize(valueBytes)
	return value, nil
}

// Set 根据key，value将目标对象存入缓存中
func (c Cache) Set(key string, value interface{}) {
	valueBytes := utils.Serialize(value)
	err := c.BigCache.Set(key, valueBytes)
	if err != nil {
		panic(err)
	}
}

// InitBigCacheConfig 初始化BigCache
func InitBigCacheConfig() {
	log.Info("初始化缓存…… BigCache")
	config := bigcache.Config{
		Shards:           1024,                      // 存储的条目数量，值必须是2的幂
		LifeWindow:       math.MaxInt16 * time.Hour, // 超时后条目被处理
		CleanWindow:      2 * time.Minute,           // 处理超时条目的时间范围
		MaxEntrySize:     500,                       // 条目最大尺寸，以字节为单位
		HardMaxCacheSize: 0,                         // 设置缓存最大值，以MB为单位，超过了不在分配内存。0表示无限制分配
	}
	bigCache, err := bigcache.NewBigCache(config)
	if err != nil {
		panic(fmt.Errorf("初始化BigCache: %s \n", err))
	}
	BigCache = &Cache{
		BigCache: bigCache,
	}
	log.Info("BigCache: 初始化完成")
}
