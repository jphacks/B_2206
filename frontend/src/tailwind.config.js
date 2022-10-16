/** @type {import('tailwindcss').Config} */
module.exports = {
  tailwindConfig: './styles/tailwind.config.js',
  content: [
    './pages/**/*.{js,ts,jsx,tsx}',
    './components/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {},
  },
  plugins: [],
}
