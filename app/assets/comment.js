window.onload = function () {
    updateHeight();
}

function updateHeight() {
    const bodyHeight = document.body.offsetHeight;
    const footerHeight = document.querySelector('footer').offsetHeight;
    const height = bodyHeight + footerHeight + 30;
    const data = { event: 'resize', offsetHeight: height };
    // console.log("updateHeight:", height);
    window.parent.postMessage(data, '*');
}

htmx.on("htmx:afterSwap", function(evt) {
    updateHeight();
});