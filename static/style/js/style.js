 // 首页轮播
 var Mswiper = new Swiper('.m-swiper .swiper-container', {
    loop: true,
    // 如果需要分页器
    pagination: {
        clickable: true,
        el: '.m-swiper .swiper-pagination',
    },
    autoplay: {
        delay: 3000,
        stopOnLastSlide: false,
        disableOnInteraction: true,
    },
})
// 只有一个就销毁
if (Mswiper.slides.length <= 3) {
    Mswiper.destroy(false);
}