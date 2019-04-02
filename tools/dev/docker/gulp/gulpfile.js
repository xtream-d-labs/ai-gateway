"use strict";

var gulp = require('gulp');
var shell = require('gulp-shell');
var autoprefixer = require("gulp-autoprefixer");
var plumber = require("gulp-plumber");
var rename = require('gulp-rename');
var sass = require("gulp-sass");
var uglify = require("gulp-uglify");
var eslint = require('gulp-eslint');

gulp.task("api", function() {
  gulp.src("/monitor/generated/cmd/scaleshift-server/main.go")
      .pipe(shell(['docker restart api'], {ignoreErrors: true}));
});
gulp.task("web", function() {
  gulp.src("/monitor/web/package.json")
      .pipe(shell(['docker restart web'], {ignoreErrors: true}));
});

gulp.task("js", function() {
  gulp.src('/monitor/web/static/js/**/*.vue')
      .pipe(eslint('/gulp/eslint.json'))
      .pipe(eslint.formatEach('compact', process.stderr))
      .pipe(rename({extname: '.js'}))
      .pipe(uglify())
      .pipe(gulp.dest("/monitor/web/static/js/min"));
});
gulp.task("sass", function() {
  gulp.src("/monitor/web/static/scss/**/*.scss")
      .pipe(plumber())
      .pipe(sass())
      .pipe(autoprefixer())
      .pipe(gulp.dest("/monitor/web/static/css"));
});

gulp.task("default", function() {
    gulp.watch("/monitor/controllers/*.go", ["api"]);
    gulp.watch("/monitor/lib/*.go",         ["api"]);
    gulp.watch("/monitor/rescale/*.go",     ["api"]);
    gulp.watch("/monitor/web/content/**/*.md",    ["web"]);
    // gulp.watch("/monitor/web/layouts/**/*.html",  ["web"]);
    gulp.watch("/monitor/web/config.toml",        ["web"]);
    gulp.watch("/monitor/web/static/js/**/*.vue", ["js"]).on("change", function(file) {
      gulp.src(file.path).pipe(shell(['docker restart web'], {ignoreErrors: true}));
    });
    gulp.watch("/monitor/web/static/scss/**/*.scss", ["sass"]);
});
