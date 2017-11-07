all: js css font html

.PHONY: js
js:
	mkdir -p dist/assets/js/
	cp -r node_modules/vue/dist/vue.js dist/assets/js/
	cp -r node_modules/element-ui/lib/index.js dist/assets/js/element.js
	gopherjs build pkg/login.go pkg/config.go -o dist/assets/js/login.js


.PHONY: css
css:
	mkdir -p dist/assets/css/
	sassc -t compressed src/scss/style.scss dist/assets/css/style.css

.PHONY: font
font:
	mkdir -p dist/assets/fonts/
	cp -r node_modules/element-theme-chalk/src/fonts/* dist/assets/fonts/
	cp -r node_modules/font-awesome/fonts/* dist/assets/fonts/

.PHONY: html
html:
	mkdir -p dist/
	cp -r src/html/* dist/
