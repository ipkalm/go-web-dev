"use strict";

window.addEventListener("DOMContentLoaded", () => {
    localStorage.setItem('favFlav','vanilla');
    let taste = localStorage.getItem('favFlav');
    console.log(taste);

    // localStorage.removeItem('favFlav');
    taste = localStorage.getItem('favFlav');
    console.log(taste);
});