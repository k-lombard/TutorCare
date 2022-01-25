const { addTailwindPlugin } = require("@ngneat/tailwind");
const tailwindConfig = require("./tailwind.config.js");

// module.exports = (config) => {
//   addTailwindPlugin({
//     webpackConfig: config,
//     tailwindConfig,
//     patchComponentsStyles: true
//   });
//   return config;
// };

module.exports = {
  module: {
    rules: [
      {
        test: /tailwind\.scss$/,
        loader: "postcss-loader",
        options: {
          postcssOptions: {
            ident: "postcss",
            syntax: "postcss-scss",
            plugins: [
              require("postcss-import"),
              require("tailwindcss"),
              require("autoprefixer"),
            ],
          },
        },
      },
    ],
  },
};
