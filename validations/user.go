package validations

type Register struct {
	Phone   string `valid:"Required;Match(/^1[34578]\\d{9}$/)"`
	Newpwd  string `valid:"Required;Match(/^[A-Za-z0-9]{6,16}$/)"`
	Smsimg  string `valid:"Required"`
	Smscode string `valid:"Required;Numeric"`
}

func RegisterMsg() map[string]string {
	arr := make(map[string]string)
	arr["Phone.Required"] = "手机号必填"
	arr["Phone.Match"] = "手机号格式不格式"
	arr["Newpwd.Required"] = "密码必填"
	arr["Newpwd.Match"] = "密码格式不符合格式"
	arr["Smsimg.Required"] = "图形验证码必填"
	arr["Smscode.Numeric"] = "数字验证码必填"
	arr["Smscode.Required"] = "短信验证码必填"
	return arr
}
