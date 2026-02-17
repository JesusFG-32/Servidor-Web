console.log("Client-side script loaded.");

document.addEventListener('DOMContentLoaded', function () {
    const ua = navigator.userAgent;
    let browserName = "Desconocido";
    let osName = "Desconocido";
    if (ua.includes("Firefox")) browserName = "Mozilla Firefox";
    else if (ua.includes("Chrome")) browserName = "Google Chrome";
    else if (ua.includes("Safari")) browserName = "Apple Safari";

    if (ua.includes("Win")) osName = "Windows";
    else if (ua.includes("Mac")) osName = "macOS";
    else if (ua.includes("Linux")) osName = "Linux";
    else if (ua.includes("Android")) osName = "Android";
    else if (ua.includes("like Mac")) osName = "iOS";

    const browserEl = document.getElementById('browser');
    const osEl = document.getElementById('os');

    if (browserEl) browserEl.innerText = browserName;
    if (osEl) osEl.innerText = osName;
});
