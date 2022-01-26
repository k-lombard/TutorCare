module.exports = {
  prefix: '',
  purge: {
    content: [ './src/**/*.{html,ts}' ],
  },
  darkMode: 'class',
  theme: {
    screens: {
      'tablet': '640px',
      // => @media (min-width: 640px) { ... }

      'laptop': '1024px',
      // => @media (min-width: 1024px) { ... }

      'desktop': '2236px',
      // => @media (min-width: 1280px) { ... }
    },
  },
    variants: {
      extend: {},
    },
    plugins: [require('@tailwindcss/forms'),require('@tailwindcss/typography')],
}
