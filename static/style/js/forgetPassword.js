var $getCode = $('#smsVerifyBtn');

var dialog = YDUI.dialog;
var phone = $("#username");
var newpwd = $("#newUserpwd");
var firmpwd = $("#firmUserpwd");
var smsimg = $("#smsGrap");
var smscode = $("#smsVerify");


var telReg = /^1[34578]\d{9}$/;
$("#reginsterBtn").click(function () {
    if (phone.val() == "") {
        dialog.toast('手机号不能为空', 'none', 1000);
        return false;
    } else if (!(telReg.test(phone.val()))) {
        dialog.toast('手机号格式不正确', 'none', 1000);
        return false;
    } else if (newpwd.val() == "") {
        dialog.toast('密码不能为空', 'none', 1000);
        return false;
    } else if (newpwd.val().length < 6) {
        dialog.toast('至少6位数字或字母', 'none', 1000);
        return false;
    }  else if (firmpwd.val() != newpwd.val()) {
        dialog.toast('两次密码不一致', 'none', 1000);
        return false;
    } else if (smsimg.val() == "") {
        dialog.toast('图片验证码不能为空', 'none', 1000);
        return false;
    } else if (smsimg.val() != "123456") {
        dialog.toast('图片验证码错误', 'none', 1000);
        return false;
    } else if (smscode.val() == "") {
        dialog.toast('短信验证码不能为空', 'none', 1000);
        return false;
    } else if (smscode.val() != "123456") {
        dialog.toast('短信验证码错误', 'none', 1000);
        return false;
    } else {
        dialog.toast('修改成功', 'none', function () {
            /* 关闭后调用 */
        });
        return true;
    }
})

/* 定义参数 */
$getCode.sendCode({
    disClass: 'btn-disabled',
    secs: 60,
    run: false,
    runStr: '{%s}秒后重新获取',
    resetStr: '重新获取验证码'
});

$getCode.on('click', function () {
    /* ajax 成功发送验证码后调用【start】 */
    YDUI.dialog.loading.open('发送中');
    setTimeout(function () {
        YDUI.dialog.loading.close();
        $getCode.sendCode('start');
        YDUI.dialog.toast('已发送', 'success', 1000);
    }, 1000);
});