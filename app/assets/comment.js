window.onload = function () {
    const bodyHeight = document.body.offsetHeight;
    const footerHeight = document.querySelector('footer').offsetHeight;
    const data = { event: 'resize', offsetHeight: bodyHeight + footerHeight };
    window.parent.postMessage(data, '*');
}