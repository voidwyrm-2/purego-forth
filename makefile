all: init
	gcc -shared -o include/stdlib.so stdlib/*.c

one: init
	gcc -shared -o include/$(LIB).so stdlib/$(LIB).c

init: clean
	mkdir -p include

clean:
	rm -rf include
