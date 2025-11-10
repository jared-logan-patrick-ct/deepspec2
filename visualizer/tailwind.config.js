/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    extend: {
      colors: {
        ct: {
          yellow: '#FFC806',
          green: '#13C1C1',
          purple: '#6259FE',
        },
        bg: {
          darkest: '#0a0a0a',
          dark: '#111',
          medium: '#1a1a1a',
          light: '#2a2a2a',
        },
        text: {
          white: '#ffffff',
          gray: '#9ca3af',
        },
        border: {
          dark: '#1a1a1a',
          medium: '#2a2a2a',
        },
      },
    },
  },
  plugins: [],
};
