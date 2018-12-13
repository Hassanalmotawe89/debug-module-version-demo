package main

import (
	"fmt"
	"runtime/debug"

	_ "rsc.io/quote"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		panic("couldn't read build info")
	}

	fmt.Printf("%s version %s\n", bi.Path, bi.Main.Version)

	for _, d := range bi.Deps {
		fmt.Printf("\tbuilt with %s version %s\n", d.Path, d.Version)
	}
}
