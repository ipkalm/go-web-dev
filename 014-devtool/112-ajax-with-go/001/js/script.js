"use strict";

window.addEventListener("DOMContentLoaded", () => {
    const makeRequest = () => {
        const xhr = new XMLHttpRequest();
        xhr.open("GET", "/boo", true)
        xhr.onreadystatechange = () => {
            if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
                alert(xhr.responseText);
            }
        }
        xhr.send();
    }

    document.querySelector("h1").onclick = makeRequest;
});