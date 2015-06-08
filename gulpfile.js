// Imports

var gulp   = require('gulp');
var less   = require('gulp-less');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');
var minify = require('gulp-minify-css');

// Less Task

gulp.task('less', function() {
	return gulp.src('./resources/less/main.less')
		.pipe(less())
		.pipe(concat('style.css'))
		.pipe(minify())
		.pipe(gulp.dest('./public/assets/css'));
});

// CSS Task

gulp.task('css', function() {
	return gulp.src([
			// Bower CSS Files

			'normalize.css/normalize.css',

		].map(function(str) { return './bower_components/' + str }))
		.pipe(concat('libraries.css'))
		.pipe(minify())
		.pipe(gulp.dest('./public/assets/css'));
});

// JavaScript Task

gulp.task('js', function() {
	return gulp.src([
			// Bower JavaScript Files

			// TODO

		].map(function(str) { return './bower_components/' + str }).concat([
			// Asset JavaScript Files

			'main.js',

		].map(function(str) { return './resources/js/' + str })))
		.pipe(concat('scripts.js'))
		.pipe(uglify())
		.pipe(gulp.dest('./public/assets/javascript'));
});

// Watch Task

gulp.task('watch', function() {
	var watch   = require('gulp-watch');

	gulp.watch('resources/js/*.js', ['js']);
	gulp.watch('resources/less/*.less', ['less']);
});

// Default Task

gulp.task('default', ['less', 'css', 'js']);
