window.addEventListener("scroll", () => {
  const maBarre = document.querySelector("header");
  if (window.scrollY > 20) {
    maBarre.classList.add("scroll-shadow");
  } else {
    maBarre.classList.remove("scroll-shadow");
  }
});
