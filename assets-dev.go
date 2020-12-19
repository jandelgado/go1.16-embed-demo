// +build dev
// access assets through the filesystem.

package main

import "os"

// in a dev build (go build -tag dev ...), access the assets using the filesystem
// assumes that binary is run in parent dirctory of assets folder.
var assets = os.DirFS(".")
