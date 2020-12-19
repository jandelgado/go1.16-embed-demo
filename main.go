// go 1.16 "embed.FS" demo
// (c) Jan Delgado 12-2020
package main

import (
	"fmt"
	"io/fs"
)

func printDir(assets fs.FS, name string) {
	files, err := fs.ReadDir(assets, name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("contents of dir %s\n", name)
	for _, entry := range files {
		fmt.Printf("  name=%-10s isdir=%v\n", entry.Name(), entry.IsDir())
	}
}

func printFile(assets fs.FS, name string) {
	contents, err := fs.ReadFile(assets, name)
	if err != nil {
		panic(err)
	}

	fmt.Printf("file '%s': %s \n", name, contents)
}

func main() {
	printDir(assets, "assets")
	printFile(assets, "assets/file1")
	printFile(assets, "assets/subdir/file3")
}
