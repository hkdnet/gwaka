package main

import (
	"fmt"
	"os"

	"github.com/hkdnet/gwaka"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Error:\n%s\n", err)
			os.Exit(1)
		}
	}()
	os.Exit(_main())
}

func _main() int {
	cli := gwaka.Gwaka{}
	cli.Parse()
	return 0
}
