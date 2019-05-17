
var dialog = YDUI.dialog;
var $ressUser = $("#ressUser"),
    $ressTel = $("#ressTel"),
    $ressPcd = $("#ressPcd"),
    $ressAddress = $("#ressAddress"),
    $ressBtn = $("#ressBtn"),
    $telReg = /^1[34578]\d{9}$/,
    $nameReg = /^[\u4E00-\u9FA5]{2,4}$/;

$ressBtn.click(function () {
    if ($ressUser.val() == "") {
        dialog.toast('请输入收货人姓名', 'none', 1000);
        return false;
    } else if (!($nameReg.test($ressUser.val()))) {
        dialog.toast(' 收货人姓名不能少于2个汉字', 'none', 1000);
        return false;
    } else if (!($telReg.test($ressTel.val()))) {
        dialog.toast('请输入11位手机号码', 'none', 1000);
        return false;
    } else if ($ressPcd.val() == "") {
        dialog.toast('请选择省份', 'none', 1000);
        return false;
    } else if ($ressAddress.val().length < 5) {
        dialog.toast('详细地址不能少于5个字，不能多于120个字', 'none', 1000);
        return false;
    } else {
        console.log("修改成功")
        return true;
    }
})

$ressPcd.citySelect();
$ressPcd.on('click', function () {
    $ressPcd.citySelect('open');
});
$ressPcd.on('done.ydui.cityselect', function (ret) {
    $(this).val(ret.provance + ' ' + ret.city + ' ' + ret.area);
});