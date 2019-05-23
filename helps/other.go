package helps

import (
	"encoding/json"
	"math/rand"
	"time"
)

// 随机一个范围的数字
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

// json数据转化为切片
func JsonToSlice(data string) (bool, []string) {
	var a []string
	dataByte := []byte(data)
	err := json.Unmarshal(dataByte, &a)
	if err == nil {
		return true, a
	}
	return false, a
}
