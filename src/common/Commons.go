package common

import (
	"math/rand"
	"strconv"
	"time"
)

//生成数据库主键
func GenId() int {
	rand.Seed(time.Now().UnixNano())
	s := strconv.Itoa(rand.Intn(10000)) + strconv.Itoa(int(time.Now().Unix()))
	id, _ := strconv.Atoi(s)
	return id
}
