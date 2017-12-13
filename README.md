![rehover](https://hoverguys.github.io/rehover/images/logo.svg)
<a href="https://rehover-build.ovo.ovh/docs/"><img src="https://codedocs.xyz/doxygen/doxygen.svg"/></a>

A homebrew racing game for the Nintendo Gamecube and Wii inspired by Diddy Kong Racing.

**Currently under heavy development**

## Compiling

You will need:
- A healthy fear of the end
- devkitPro (with devkitPPC)
- Go 1.8+
- Cmake 3.1+
- [ppc-portlibs](https://github.com/Hamcha/ppc-portlibs) libraries installed, specifically:
    - entityx

#### Step 0: BUILD TOOLS
Run `build.sh` or `build.cmd` (depending on your OS of choice) in the `tools` folder

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
