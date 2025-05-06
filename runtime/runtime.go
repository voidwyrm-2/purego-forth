package runtime

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/voidwyrm-2/purego-forth/ccall"
)

type ExternFunc func(func(int), func() int, int) int

type Interpreter struct {
	libs map[string]ccall.CCaller
	env  *ccall.CCaller
}

func New(preload []string) (Interpreter, error) {
	libs := map[string]ccall.CCaller{}

	for _, p := range preload {
		p = filepath.Clean(p)

		caller, err := ccall.New(p)
		if err != nil {
			return Interpreter{}, err
		}

		libs[filepath.Base(p)] = caller
	}

	var env *ccall.CCaller
	if len(preload) > 0 {
		l := libs[filepath.Base(filepath.Clean(preload[0]))]
		env = &l
	}

	return Interpreter{libs: libs, env: env}, nil
}

func (i *Interpreter) Execute(text string) ([]int, error) {
	tokens := strings.Fields(text)

	pos := 0

	errf := func(str string, a ...any) ([]int, error) {
		return []int{}, fmt.Errorf(str, a...)
	}

	stack := []int{}

	push := func(i int) {
		stack = append(stack, i)
	}

	pop := func() int {
		value := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return value
	}

	for pos < len(tokens) {
		tok := tokens[pos]

		if n, err := strconv.Atoi(tok); err == nil {
			push(n)
			pos++
		} else if tok == "echo" { // have to implement echo and emit this way because printf won't work for some reason
			if len(stack) == 0 {
				return errf("stack underflow")
			}

			fmt.Println(pop())
			pos++
		} else if tok == "emit" {
			if len(stack) == 0 {
				return errf("stack underflow")
			}

			fmt.Println(rune(pop()))
			pos++
		} else if tok == "dump" {
			fmt.Println(stack)
			pos++
		} else {
			var env *ccall.CCaller

			if strings.Contains(tok, ":") {
				p := strings.Split(tok, ":")
				if lib, ok := i.libs[p[0]]; !ok {
					return errf("library '%s' does not exist")
				} else {
					env = &lib
				}
			} else {
				if i.env == nil {
					return errf("env is not available")
				}

				env = i.env
			}

			var fptr ExternFunc

			env.Retrieve(tok, &fptr)

			switch rc := fptr(push, pop, len(stack)); rc {
			case 0:
				pos++
			case 1:
				return errf("stack underflow")
			default:
				return errf("invalid return code %d", rc)
			}
		}
	}

	return stack, nil
}
