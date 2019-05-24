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
