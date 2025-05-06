LIB=include/stdlib.so

all: init build demo

build:
	gcc -shared -o $(LIB) stdlib/*.c

one: init
	gcc -shared -o $(LIB) stdlib/$(LIB).c

init: clean
	mkdir -p include

clean:
	rm -rf include

demo:
	go run . -f examples/math.pfth $(LIB)
	go run . -f examples/hello.pfth $(LIB)
