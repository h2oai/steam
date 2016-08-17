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
    path: path.join(__dirname, '../var/master/www'),
    publicPath: '/',
    filename: 'bundle.js'
  },
  resolve: {
    extensions: ['', '.ts', '.tsx', '.js', '.css']
  },
  module: {
    preLoaders: [
      {
        test: /\.tsx?$/,
        exclude: [path.resolve(__dirname, 'src/Proxy/CLI.ts'), path.resolve(__dirname, 'src/Proxy/Proxy.ts'), path.resolve(__dirname, 'src/Proxy/xhr.ts')],
        loader: 'tslint'
      }
    ],
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
  sassResources: [],
  externals: {
    'cheerio': 'window',
    'react/addons': true,
    'react/lib/ExecutionEnvironment': true,
    'react/lib/ReactContext': true
  },
  plugins: [
    new BrowserSyncPlugin({
      host: 'localhost',
      port: 3000,
      server: { baseDir: ['../var/master/www'] }
    }),
    require('webpack-fail-plugin')
  ],
  tslint: {
    configuration: {
      rules: {
        'triple-equals': [true, 'allow-null-check', 'allow-undefined-check'],
        'class-name': true,
        'semicolon': [true, 'always', 'ignore-interfaces'],
        'radix': true,
        'align': [true, 'parameters', 'statements'],
        'variable-name': [true, 'ban-keywords', 'check-format', 'allow-leading-underscore'],
        'whitespace': [true, 'check-branch', 'check-decl', 'check-operator', 'check-module', 'check-separator', 'check-type', 'check-typecast'],
        'no-console': true,
        'no-construct': true,
        'curly': true,
        'no-eval': true,
        'no-duplicate-key': true,
        indent: 'spaces',
        'linebreak-style': [true, 'LF'],
        'quotemark': [true, 'single', 'jsx-double'],
        'one-variable-per-declaration': [true, 'ignore-for-loop'],
        'no-angle-bracket-type-assertion': true,
        'switch-default': true,
        'no-conditional-assignment': true,
        'no-bitwise': true,
        'no-debugger': true,
        'forin': true,
        'no-duplicate-variable': true,
        'no-invalid-this': true,
        'no-string-literal': true
      }
    },
    emitErrors: true,
    failOnHint: true
  }
};
