// @ts-check
const navSelector = '.top-nav a';
const activeClass = 'active';

/**
 * Initializes the top navigation links.
 */
function initializeNavLinks() {
  const currentPath = window.location.pathname;
  const rootPath = '/' + currentPath.split('/')[1];

  /** @type {NodeListOf<HTMLLinkElement>} */
  const links = document.querySelectorAll(navSelector);

  links.forEach((link) => {
    const linkPath = new URL(link.href).pathname;

    if (linkPath === rootPath && rootPath !== '/') {
      link.classList.add(activeClass);
    } else {
      link.classList.remove(activeClass);
    }
  });
}

initializeNavLinks();
