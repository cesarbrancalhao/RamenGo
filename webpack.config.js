const path = require('path');
const Dotenv = require('dotenv-webpack');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');

module.exports = {
  entry: './client/index.js',
  output: {
    filename: 'main.js',
    path: path.resolve(__dirname, './client/src/scripts/dist'),
  },
  plugins: [
    new Dotenv(),
    new CleanWebpackPlugin(),
  ],
};