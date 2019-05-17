package validations

import "github.com/astaxie/beego/validation"

type Register struct {
	Phone   string `valid:"Required;Match(/^1[34578]\\d{9}$/)"`
	Newpwd  string `valid:"Required;Match(/^[A-Za-z0-9]{6,16}$/)"`
	Smsimg  string `valid:"Required"`
	Smscode string `valid:"Required;Numeric"`
}

func Prompt() {
	v := validation.Validation{}
	v.SetError("phone.Required", "测试")
	v.SetError("phone.Match", "测试22")
}
