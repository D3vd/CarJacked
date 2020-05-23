module.exports = {
  plugins: [
    "gatsby-plugin-sass",
    "gatsby-plugin-react-helmet",
    {
      resolve: `gatsby-plugin-create-client-paths`,
      options: { prefixes: [`/case/*`] },
    },
  ],
};
