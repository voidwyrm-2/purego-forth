package main

import (
	"flag"
	"io"
	"os"

	"github.com/voidwyrm-2/purego-forth/runtime"
)

func _main() error {
	inputPath := flag.String("f", "", "The file to interpret")
	readStdin := flag.Bool("in", false, "Read from stdin")

	flag.Parse()

	libraries := flag.Args()

	content := ""

	if *readStdin {
		raw, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}

		content = string(raw)
	} else {
		raw, err := os.ReadFile(*inputPath)
		if err != nil {
			return err
		}

		content = string(raw)
	}

	interp, err := runtime.New(libraries)
	if err != nil {
		return err
	}

	_, err = interp.Execute(content)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := _main(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
