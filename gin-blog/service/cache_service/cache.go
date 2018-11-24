package cache_service

import (
	"fmt"
	"test.go.dev/gin-blog/pkg/gredis"
)

//初始化redis
func init() {
	if err := gredis.Setup(); err != nil {
		fmt.Println(err)
	}

}
