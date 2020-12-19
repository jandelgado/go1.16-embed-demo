# Go 1.16 embed.FS and fs.FS demo

Short demo of upcoming go 1.16 "io/fs" package and file embedding feature.

Depending on how the binary is built, either with `dev` build tag or without,
the `assets/` folder will be embedded in the code (without `dev` tag), or the
assets will be read from the filesystem.

Run `make` to build the demo. The `embed-demo` will have the assets folder 
compiled-in. The `embed-demo-dev` will use the filesystem to obtain the assets.

See also 
* https://tip.golang.org/doc/go1.16#library-embed
* https://go.googlesource.com/proposal/+/master/design/draft-embed.md
* https://go.googlesource.com/proposal/+/master/design/draft-iofs.md

Copyright 2020 by Jan Delgado, License MIT.


