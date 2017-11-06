# rehover

A rewrite of that stupid [hovercraft gamecube homebrew](https://github.com/hamcha/hovercraft).

## Compiling

You will need:
- A healthy fear of the end
- devkitPro (w/ devkitPPC)
- Cmake

#### Step 0: BUILD TOOLS
Run `build.sh` or `build.cmd` depending on your OS of choice

#### Step 1: BUILD THE PROJECT
```
mkdir build && cd build
cmake -DCMAKE_BUILD_TYPE=Release ../project
make
```

**Note:** Build type DEBUG is currently broken, if you get a blank screen make sure `CMAKE_BUILD_TYPE` is set to `Release`

#### Step 2:
There is no step 2, use the compiled `rehover.dol` with your favorite emulator (if it's not [Dolphin](https://dolphin-emu.org/) we need to talk)
