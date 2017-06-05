
CSS_SRCS = $(wildcard resources/scss/*.scss)
OBJ = $(CSS_SRCS:scss=css)

GOPHERJS_SRC = $(wildcard client/*.go)


resources/scss/%.css : resources/scss/%.scss
	sass $< $@

all: $(OBJ) resources/main.js

resources/main.js: $(GOPHERJS_SRC)
	gopherjs build github.com/kori-irrlicht/artificial-horizon/client -o resources/main.js -m
