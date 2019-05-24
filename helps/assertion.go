package helps

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

// 断言转化为float64形式
func AssertionToFloat64(s interface{}) float64 {
	value, ok := s.(float64)
	if ok {
		return value
	}
	return 0.00
}
