/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/**/*.html"],
  theme: {
    extend: {
      colors: {

        heading: "var(--heading-color)",
        text: {
          DEFAULT: "var(--text-color)", post: {
            sm: "var(--post-sm-color)",
            lg: "var(--post-heading-color)",
            tags: "var(--post-tags-text-color)"
          },
          nav: {
            logo: "var(--nav-logo-color)",
            logo2: "var(--nav-logo-2-color)"
          },
        },
        bg: {
          DEFAULT: "var(--bg-color)",
          post: "var(--post-bg-color)",
          tags: "var(--tags-box-bg-color)"
        },
      },
    },
  },
  plugins: [],
}
