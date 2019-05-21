package common

import (
	"math/rand"
	"time"
)

//生成数据库主键
func GenId() int64 {
	rand.Seed(time.Now().UnixNano())
	s := rand.Intn(10000) + int(time.Now().Unix())
	return int64(s)
}
