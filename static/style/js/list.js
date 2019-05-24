$(".m-screen .screen-item").click(function () {
    $(this).addClass("active").siblings().removeClass("active");
});