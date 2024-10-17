const themeBtn = document.getElementById('theme-btn');
const svgMoon = document.getElementById('svg-moon');
const svgSun = document.getElementById('svg-sun');
const themeLink = document.getElementById('hljs-theme');

// Function to update the theme link
function updateCodeTheme(theme) {
   console.log('🚀  theme :- ', theme)
   if (theme === 'dark') {
      themeLink.href = `https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.9.0/build/styles/a11y-dark.min.css`;
   } else {
      themeLink.href = `https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.9.0/build/styles/a11y-light.min.css`;
   }
}

// Check for dark mode preference
const prefersDarkScheme = window.matchMedia && window.matchMedia("(prefers-color-scheme: dark)").matches;

// Function to change the theme
function changeTheme(theme) {
   const isDark = theme === 'dark';

   updateCodeTheme(theme);

   document.documentElement.setAttribute('data-theme', theme);
   svgMoon.style.display = isDark ? 'none' : 'block';
   svgSun.style.display = isDark ? 'block' : 'none';
}

// Toggle theme on button click
themeBtn.addEventListener('click', () => {
   const currentTheme = document.documentElement.getAttribute('data-theme') || (prefersDarkScheme ? 'dark' : 'light');
   changeTheme(currentTheme === 'dark' ? 'light' : 'dark');
});

// Set initial theme on page load
window.onload = () => {
   const initialTheme = prefersDarkScheme ? 'dark' : 'light';
   changeTheme(initialTheme);
};
