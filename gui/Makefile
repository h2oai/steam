.PHONY: \
	setup \
	all \
	js \
	html \
	css \
	lib \
	test

O=bin
N=./node_modules/.bin
B=./bower_components

all: setup lib js html css

setup:
	npm install
	$N/bower install
	$N/tsd install
	mkdir -p $O

lib:
	rsync -u $B/jquery/dist/jquery.min.js $O/
	rsync -u $B/async/dist/async.min.js $O/
	rsync -u $B/underscore/underscore-min.js $O/
	rsync -u $B/knockout/dist/knockout.js $O/
	rsync -u $B/mousetrap/mousetrap.min.js $O/
	rsync -u $B/faker/build/build/faker.min.js $O/
	rsync -u $B/HTML5-Reset/assets/css/reset.css $O/
	rsync -u $B/qunit/qunit/qunit.css $O/
	rsync -u $B/qunit/qunit/qunit.js $O/
	mkdir -p $O/fonts
	rsync -u tmpl/favicon.ico $O/
	$N/jade --pretty --out $O ./tmpl/test.jade

js:
	$N/tsc -p src

html:
	cat tmpl/template.jade | perl -pe 's@^    template\.(\S+)@    script(type="text/html" id="tmpl-$$1")@g' > tmpl/index.jade
	$N/jade --pretty --out $O tmpl/index.jade

css:
	$N/stylus --use ./node_modules/nib/lib/nib --out $O tmpl/styles.styl

test:
	$N/tsc -p test
	$N/karma start karma-conf.js --single-run


