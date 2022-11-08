package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gosimple/slug"
)

// slugtext "asd qwe 923864ľíščýáž5ť%" -> "asd-qwe-923864liscyaz5t"
func slugtext() {
	if len(os.Args) > 1 {
		arg := os.Args[1:]
		text := slug.Make(strings.Join(arg, " "))
		fmt.Println(text)
	} else {
		fmt.Printf("Usage:\v%v <text-to-slug>\n", os.Args[0])
		os.Exit(2)
	}
}

func main() {
	slugtext()
}
