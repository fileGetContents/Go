package helps

import (
	"math/rand"
	"time"
)

func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

//进行断言操作
func NilToString(s interface{}) string {
	value, ok := s.(string)
	if ok {
		return value
	}
	return ""
}

// 断言转化为Int形式
func AssertionToInt(s interface{}) int {
	value, ok := s.(int)
	if ok {
		return value
	}
	return -1
}
