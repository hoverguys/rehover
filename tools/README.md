# rehover build tools

This folder contains build tools required to be able to build rehover

## Building

Use either `build.cmd` or `build.sh`, depending on your OS of choice, to build all the tools in a `bin/` folder. Once compiled and correctly positioned by the build script, Cmake should have no problem finding them.

## Tools

### bento

bento is an alternative to devkitpro's bin2o in Go to generate the assembly file and header from a generic resource in one single binary instead of relying on makefiles. It also makes it easier to customize both outputs via Go's text/template library.

### objconv

objconv is a rewrite of our previous [obj2bin](https://github.com/Hamcha/hovercraft/tree/master/tools/obj2bin_src) tool to generate BMB files from OBJ models so that it doesn't require big dependencies such as Assimp and Boost.

### texconv

texconv is an alternative to [gxtool/gxtexconv](https://github.com/devkitPro/gamecube-tools/tree/master/gxtool) in Go that converts texture to our own texture format (BTB) instead of dkP's TPL.

### gcpacker

gcpacker is a tool that generates a single GCR (GameCube Resource) file from a list of resources (with identifier and paths) provided as a line-separated text file.