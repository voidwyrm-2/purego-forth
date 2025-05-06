package main

import (
	"flag"
	"os"

	"github.com/voidwyrm-2/purego-forth/runtime"
)

func _main() error {
	inputPath := flag.String("f", "", "The file to interpret")

	flag.Parse()

	libraries := flag.Args()

	content, err := os.ReadFile(*inputPath)
	if err != nil {
		return err
	}

	interp, err := runtime.New(libraries)
	if err != nil {
		return err
	}

	_, err = interp.Execute(string(content))
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
