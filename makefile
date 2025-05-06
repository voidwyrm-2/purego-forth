all: build demo

build: init
	# -fPIC needed or stdout/stderr cause link error
	gcc -shared -fPIC -o include/stdlib.so stdlib/*.c

init: clean
	mkdir -p include

clean:
	rm -rf include

demo:
	make build

	go run . -f examples/hello.pfth include/stdlib.so
	go run . -f examples/math.pfth include/stdlib.so
