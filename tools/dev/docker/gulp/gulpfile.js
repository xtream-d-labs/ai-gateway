"use strict";

var gulp = require('gulp');
var shell = require('gulp-shell');
var autoprefixer = require("gulp-autoprefixer");
var plumber = require("gulp-plumber");
var rename = require('gulp-rename');
var sass = require("gulp-sass");
var uglify = require("gulp-uglify");
var eslint = require('gulp-eslint');

gulp.task("js", function() {
  gulp.src('/monitor/web/src/static/js/**/*.vue')
      .pipe(eslint('/gulp/eslint.json'))
      .pipe(eslint.formatEach('compact', process.stderr))
      .pipe(rename({extname: '.js'}))
      .pipe(uglify())
      .pipe(gulp.dest("/monitor/web/src/static/js/min"));
});
gulp.task("sass", function() {
  gulp.src("/monitor/web/src/static/scss/**/*.scss")
      .pipe(plumber())
      .pipe(sass())
      .pipe(autoprefixer())
      .pipe(gulp.dest("/monitor/web/src/static/css"));
});

gulp.task("default", function() {
  gulp.watch("/monitor/web/src/static/js/**/*.vue", ["js"]).on("change", function(file) {
    gulp.src(file.path).pipe(shell(['docker restart scaleshift_web'], {ignoreErrors: true}));
  });
  gulp.watch("/monitor/web/src/static/scss/**/*.scss", ["sass"]);
});
