# Purego Forth

### This branch is for testing C++ interop, nothing is guaranteed to work

This is a Forth dialect whose standard library is (almost) entirely written in C, via ebitengine's [purego](https://github.com/ebitengine/purego)

# Installation

At the moment, the only way to install is by compiling it yourself

## Compiling Locally

Compiling locally requires both Go and GCC

On Windows, GCC can be obtained via [MSYS2](https://www.msys2.org) or WSL

1. Install Go from its [website](https://go.dev)
2. Open a terminal of your choice
3. `git clone https://github.com/voidwyrm-2/purego-forth`
4. `cd purego-forth`
5. `make build`
6. `go run . -f [file] [libraries...]`, e.g. `go run . -f examples/hello.pfth include/stdlib.so`

<!--
# Creating your own libraries

Creating your own library is very easy, you just need to know a bit of C.
-->
