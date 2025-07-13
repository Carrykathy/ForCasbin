package component

import (
	"fmt"
	"time"
)

var (
	DB          *gorm.DB
	GlobalCache *bigcache.BigCache
)

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go_blog?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(fmt.Sprintf("connect to database failed, err: %v", err))
	}

	GlobalCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(time.Minute * 10))
	if err != nil {
		panic(fmt.Sprintf("create bigcache failed, err: %v", err))
	}
}
