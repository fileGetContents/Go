package helps

import (
	"encoding/json"
)

//type Server struct {
//	ServerName string
//	ServerIP   string
//}
//
//type Serverslice struct {
//	Servers []Server
//}
//
//func Suc() {
//	var s Serverslice
//	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
//	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
//	b, err := json.Marshal(s)
//	if err != nil {
//		fmt.Println("json err:", err)
//	}
//	fmt.Println(string(b))
//}

type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Statuss struct {
	CodeNum int      `json:"code_num"`
	Statuss []Status `json:"statuss"`
}

// 成功无数据状态
func Suc(msg string) string {
	if msg == "" {
		msg = "ok"
	}
	var s Status
	s.Msg, s.Code = msg, 0
	r, _ := json.Marshal(s)
	return string(r)
}

// 系统错误
func Error(msg string) string {
	if msg == "" {
		msg = "系统异常"
	}
	var s Status
	s.Msg, s.Code = msg, 10000
	r, _ := json.Marshal(s)
	return string(r)
}
