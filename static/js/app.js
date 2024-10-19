// assets/js/app.js
var navSelector = ".top-nav a";
var activeClass = "active";
function initializeNavLinks() {
  const currentPath = window.location.pathname;
  const rootPath = "/" + currentPath.split("/")[1];
  const links = document.querySelectorAll(navSelector);
  links.forEach((link) => {
    const linkPath = new URL(link.href).pathname;
    if (linkPath === rootPath && rootPath !== "/") {
      link.classList.add(activeClass);
    } else {
      link.classList.remove(activeClass);
    }
  });
}
initializeNavLinks();
//# sourceMappingURL=app.js.map
