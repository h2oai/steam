module.exports = function(config) {
  config.set({
    basePath: '',
    autoWatch: true,
    singleRun: true,
    frameworks: ['qunit'],
    files: [
      'bin/async.min.js',
      'bin/underscore-min.js',
      'bin/knockout.js',
      'bin/tests.js'
    ],
    plugins: [
      'karma-coverage',
      'karma-phantomjs-launcher',
      'karma-qunit'
    ],
    browsers: ['PhantomJS'],
    reporters: ['progress', 'coverage'],
    preprocessors: { 'bin/*.js': ['coverage'] }
  });
};
