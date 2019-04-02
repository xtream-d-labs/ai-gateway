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
  gulp.src("/monitor/api/generated/cmd/scale-shift-server/main.go")
      .pipe(shell(['docker restart api'], {ignoreErrors: true}));
});
gulp.task("web", function() {
  gulp.src("/monitor/web/src/package.json")
      .pipe(shell(['docker restart web'], {ignoreErrors: true}));
});

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
    gulp.watch("/monitor/api/auth/*.go",        ["api"]);
    gulp.watch("/monitor/api/config/*.go",      ["api"]);
    gulp.watch("/monitor/api/controllers/*.go", ["api"]);
    gulp.watch("/monitor/api/db/*.go",          ["api"]);
    gulp.watch("/monitor/api/http/*.go",        ["api"]);
    gulp.watch("/monitor/api/lib/*.go",         ["api"]);
    gulp.watch("/monitor/api/log/*.go",         ["api"]);
    gulp.watch("/monitor/api/queue/*.go",       ["api"]);
    gulp.watch("/monitor/api/reg/*.go",         ["api"]);
    gulp.watch("/monitor/api/rescale/*.go",     ["api"]);
    gulp.watch("/monitor/web/src/content/**/*.md",    ["web"]);
    // gulp.watch("/monitor/web/layouts/**/*.html",  ["web"]);
    gulp.watch("/monitor/web/src/config.toml",        ["web"]);
    gulp.watch("/monitor/web/src/static/js/**/*.vue", ["js"]).on("change", function(file) {
      gulp.src(file.path).pipe(shell(['docker restart web'], {ignoreErrors: true}));
    });
    gulp.watch("/monitor/web/src/static/scss/**/*.scss", ["sass"]);
});
