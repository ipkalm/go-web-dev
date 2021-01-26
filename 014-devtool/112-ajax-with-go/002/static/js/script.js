"use strict";

window.addEventListener("DOMContentLoaded", () => {
    const makeRequest = () => {
        const xhr = new XMLHttpRequest();
        xhr.open("GET", "/boo", true)
        xhr.onreadystatechange = () => {
            if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
                const b = document.querySelector("body");
                const myH1 = document.createElement("h1");
                const myH1Text = document.createTextNode(xhr.responseText);
                myH1.appendChild(myH1Text);
                b.appendChild(myH1);
            }
        }
        xhr.send();
    }

    document.querySelector("h1").onclick = makeRequest;
});