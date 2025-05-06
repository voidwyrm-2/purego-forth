// from https://github.com/ebitengine/purego/blob/main/examples/libc/main_windows.go

//go:build windows

package openlib

import "golang.org/x/sys/windows"

func OpenLibrary(name string) (uintptr, error) {
	dll := windows.NewLazySystemDLL(name)
	return uintptr(dll.Handle()), nil
}
