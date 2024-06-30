function comments() {
    const iframeContainer = document.getElementById("sssc-iframe-container");
    const iframe = document.createElement("iframe");

    const host = iframeContainer.dataset.host;
    const pageUrl = iframeContainer.dataset.url ? iframeContainer.dataset.url : "";
    const title = iframeContainer.dataset.title ? iframeContainer.dataset.title : "";
    const pageSlug = iframeContainer.dataset.slug;

    // Set the iframe src
    iframe.src = `${host}/comments?page_url=${pageUrl}&page_slug=${pageSlug}&title=${title}`;
    iframe.width = "90%";
    iframe.style.border = "none";
    iframe.style.borderRadius = "10px";
    iframe.style.margin = "0 auto";
    iframeContainer.appendChild(iframe);

    window.addEventListener('message', function (event) {
        if (event.data.event == 'resize') {
            // console.log("receive updateHeight", event.data.offsetHeight);
            iframe.height = event.data.offsetHeight + "px";
        }
    });
}

comments();