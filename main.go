// go 1.16 "embed" demo
// 18.12.2020 Jan Delgado
package main

import (
	"embed"
	"fmt"
	"io/fs"
)

// contents of directory "assets" are stored in variable assets
//go:embed assets
var assets embed.FS

func printDir(dir fs.ReadDirFS, name string) {
	files, err := dir.ReadDir(name)
	if err != nil {
		panic(err)
	}

	fmt.Println("contents of dir %s", name)
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
