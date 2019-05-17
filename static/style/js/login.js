var dialog = YDUI.dialog;
var phone = $("#username");
var pwd = $("#userpwd");
var telReg = /^1[34578]\d{9}$/;
$("#loginBtn").click(function () {
    if (phone.val() == "") {
        dialog.toast('手机号不能为空', 'none', 1000);
        return false;
    } else if (!(telReg.test(phone.val()))) {
        dialog.toast('账号或密码错误', 'none', 1000);
        return false;
    } else if (pwd.val() == "") {
        dialog.toast('密码不能为空', 'none', 1000);
        return false;
    } else if (pwd.val() != '123456') {
        dialog.toast('账号或密码错误', 'none', 1000);
        return false;
    } else {
        dialog.toast('登录成功', 'none', function () {
            /* 关闭后调用 */
        });
        return true;
    }
})