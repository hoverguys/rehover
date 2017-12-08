# rehover
<a href="https://rehover-build.ovo.ovh/docs/"><img src="https://codedocs.xyz/doxygen/doxygen.svg"/></a>

A rewrite of that stupid [hovercraft gamecube homebrew](https://github.com/hoverguys/hovercraft-old).

## Compiling

You will need:
- A healthy fear of the end
- devkitPro (w/ devkitPPC)
- Go 1.8+
- Cmake 3.1+
- [ppc-portlibs](https://github.com/Hamcha/ppc-portlibs) libraries installed, specifically:
    - entityx

#### Step 0: BUILD TOOLS
Run `build.sh` or `build.cmd` depending on your OS of choice

#### Step 1: BUILD THE PROJECT
```
mkdir build && cd build
cmake ../project
make
```

#### Step 2:
There is no step 2, use the compiled `rehover.dol` with your favorite emulator (if it's not [Dolphin](https://dolphin-emu.org/) we need to talk)

## License

**TL;DR**: Code is MIT, Assets are CC BY 4.0

See `LICENSE.md` for more details
