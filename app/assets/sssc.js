function comments() {
    const iframeContainer = document.getElementById("sssc-iframe-container");
    const iframe = document.createElement("iframe");

    const host = iframeContainer.dataset.host;
    const pageUrl = iframeContainer.dataset.pageUrl ? iframeContainer.dataset.pageUrl : "";
    const title = iframeContainer.dataset.title ? iframeContainer.dataset.title : "";
    const pageSlug = iframeContainer.dataset.slug;
    let height = iframeContainer.dataset.height ? iframeContainer.dataset.height : 1000;

    iframe.src = host + "/comments?page_url=" + pageUrl;
    iframe.src = `${host}/comments?page_url=${pageUrl}&page_slug=${pageSlug}&title=${title}`;
    iframe.width = "90%";
    iframe.height = height + "px";
    iframe.style.border = "none";
    iframe.style.borderRadius = "10px";
    iframe.style.margin = "0 auto";
    iframeContainer.appendChild(iframe);
}

comments();