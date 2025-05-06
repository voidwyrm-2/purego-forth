package ccall

import (
	"path/filepath"

	"github.com/ebitengine/purego"
	"github.com/voidwyrm-2/purego-forth/ccall/openlib"
)

type CCaller struct {
	handle uintptr
}

func New(path string) (CCaller, error) {
	handle, err := openlib.OpenLibrary(filepath.Clean(path))
	if err != nil {
		return CCaller{}, err
	}

	return CCaller{handle: handle}, nil
}

func (c CCaller) Retrieve(name string, fptr any) {
	purego.RegisterLibFunc(fptr, c.handle, name)
}
