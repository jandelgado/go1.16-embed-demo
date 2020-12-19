// +build !dev
// embed assets in the binary using go 1.16 embed.FS

package main

import "embed"

// when not in a dev build, contents of directory "assets" are stored in variable assets
//go:embed assets
var assets embed.FS
