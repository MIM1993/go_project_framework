package utils

import (
	"fmt"
	"math/rand"
	"time"
)

//依据时间戳产生logid，该logid方便后续根据真实时间戳分类存储图文日志
func GenLogId() string {
	return fmt.Sprintf("%d.%d", time.Now().Unix(), rand.Intn(10000))
}

