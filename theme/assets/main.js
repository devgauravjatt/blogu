console.log(" 🚀  Theme is ready to use  ");

const userPref = window.matchMedia("(prefers-color-scheme: light)").matches
	? "light"
	: "dark";
let currentTheme = localStorage.getItem("theme") ?? userPref;
const themeBtn = document.getElementById("theme-btn");
const svgMoon = document.getElementById("svg-moon");
const svgSun = document.getElementById("svg-sun");
const themeLink = document.getElementById("hljs-theme");
const goTop = document.getElementById("goTop");

// Function to change the theme
function changeTheme(theme, toggle = false) {
	console.log("🚀  theme, toggle :- ", theme, toggle);
	if (theme === "dark") {
		themeLink.href =
			"https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.9.0/build/styles/a11y-dark.min.css";
	} else {
		themeLink.href =
			"https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.9.0/build/styles/a11y-light.min.css";
	}
	document.documentElement.setAttribute("data-theme", theme);
	if (toggle) {
		localStorage.setItem("theme", theme);
	}
	svgMoon.style.display = theme === "dark" ? "none" : "block";
	svgSun.style.display = theme === "dark" ? "block" : "none";
}

// Toggle theme on button click
themeBtn.addEventListener("click", () => {
	console.log("🚀  currentTheme :- ", currentTheme);
	const newTheme = currentTheme === "dark" ? "light" : "dark";
	changeTheme(newTheme, true);
	currentTheme = newTheme; // Update the currentTheme after toggling
});

// Set initial theme on page load
window.addEventListener("DOMContentLoaded", () => {
	changeTheme(currentTheme);
});

// window listener dom reload

window.addEventListener("scroll", () => {
	if (window.scrollY > 300) {
		goTop.style.display = "block";
	} else {
		goTop.style.display = "none";
	}
});

goTop.addEventListener("click", () => {
	window.scrollTo({
		top: 0,
		behavior: "smooth",
	});
});

const currentYear = new Date().getFullYear();
document.getElementById("year").innerHTML = currentYear;
