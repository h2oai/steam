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
        loader: "tslint"
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
  sassResources: [

  ],
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
  // more options in the optional tslint object
  tslint: {
    configuration: {
      // for a complete list of rules, see: http://palantir.github.io/tslint/rules/
      rules: {
        'triple-equals': [true, 'allow-null-check', 'allow-undefined-check'],
        'class-name': true,
        'semicolon': [true, 'always', 'ignore-interfaces'],
        'radix': true,
        'align': [true, 'parameters', 'statements'],
        'variable-name': [true, 'ban-keywords', 'check-format', 'allow-leading-underscore'],
        'whitespace': [true, 'check-branch', 'check-decl', 'check-operator', 'check-module', 'check-separator', 'check-type', 'check-typecast']
        }
    },

    // tslint errors are displayed by default as warnings
    // set emitErrors to true to display them as errors
    emitErrors: false,

    // tslint does not interrupt the compilation by default
    // if you want any file with tslint errors to fail
    // set failOnHint to true
    failOnHint: false,

    // name of your formatter (optional)
    // formatter: "yourformatter",

    // path to directory containing formatter (optional)
    // formattersDirectory: "node_modules/tslint-loader/formatters/",

    // These options are useful if you want to save output to files
    // for your continuous integration server
    /*
    fileOutput: {
      // The directory where each file's report is saved
      dir: "./foo/",

      // The extension to use for each report's filename. Defaults to "txt"
      ext: "xml",

      // If true, all files are removed from the report directory at the beginning of run
      clean: true,

      // A string to include at the top of every report file.
      // Useful for some report formats.
      header: "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n<checkstyle version=\"5.7\">",

      // A string to include at the bottom of every report file.
      // Useful for some report formats.
      footer: "</checkstyle>"
    }
    */
  }
};
