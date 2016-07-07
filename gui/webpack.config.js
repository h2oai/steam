/**
 * @author justin on 6/17/16.
 */

var path = require('path');
var BrowserSyncPlugin = require('browser-sync-webpack-plugin');

module.exports = {
  entry: [
    './src/main.tsx',
    './src/index.html'
  ],
  output: {
    path: path.join(__dirname, 'dist'),
    publicPath: '/',
    filename: 'bundle.js'
  },
  resolve: {
    extensions: ['', '.ts', '.tsx', '.js', '.css']
  },
  module: {
    loaders: [
      {
        test: /\.tsx?$/,
        loader: 'ts-loader'
      },
      {
        test: /\.css$/,
        loader: 'style-loader!css-loader'
      },
      {
        test: /\.scss$/,
        loaders: ['style', 'css', 'sass']
      },
      {
        test: /\.html$|\.jpe?g$|\.gif$|\.png$|\.ico$/,
        loader: 'file?name=[name].[ext]'
      },
      {
        test: /\.(eot|woff|woff2|ttf|svg)(\?\S*)?$/,
        loader: 'url?limit=100000&name=[name].[ext]'
      },
      {
        test: /\.svg$/,
        loader: 'file-loader'
      }
    ]
  },
  sassResources: [
    
  ],
  plugins: [
    new BrowserSyncPlugin({
      host: 'localhost',
      port: 3000,
      server: { baseDir: ['dist'] }
    })
  ]
};