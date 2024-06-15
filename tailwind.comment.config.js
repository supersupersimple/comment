/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    'app/templates/comment/**/*.templ',
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('daisyui'),
  ],
  daisyui: {
    themes: ["light", "dark", "cupcake"],
  },
}

