var dialog = YDUI.dialog;
var $deleteAddress = $(".address-gnee-delete");
$deleteAddress.click(function () {
    dialog.confirm('提示', '确定删除当前地址?！', function () {
        console.log("你点了确定")
    });
})