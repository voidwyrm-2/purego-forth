all: build demo

build: init
	# -fPIC needed or stdout/stderr cause link error
	gcc -shared -fPIC -o include/_stdlibA.so stdlib/*.c
	g++ -shared -fPIC -o include/_stdlibB.so stdlib/*.cc
	gcc -shared -fPIC -o include/stdlib.so include/_stdlibA.so include/_stdlibB.so

init: clean
	mkdir -p include

clean:
	rm -rf include

demo:
	go run . -f examples/hello.pfth include/stdlib.so
	go run . -f examples/math.pfth include/stdlib.so
