// from https://github.com/ebitengine/purego/blob/main/examples/libc/main_unix.go

//go:build !windows

package openlib

import "github.com/ebitengine/purego"

func OpenLibrary(name string) (uintptr, error) {
	return purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}
