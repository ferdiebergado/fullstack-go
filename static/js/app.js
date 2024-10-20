var e=".navbar a",n="active";function s(){let a="/"+window.location.pathname.split("/")[1];document.querySelectorAll(e).forEach(t=>{new URL(t.href).pathname===a&&a!=="/"?t.classList.add(n):t.classList.remove(n)})}s();
//# sourceMappingURL=app.js.map
