![rehover](https://hoverguys.github.io/rehover/images/logo.svg)
<a href="https://rehover-build.ovo.ovh/docs/"><img src="https://codedocs.xyz/doxygen/doxygen.svg"/></a>

A homebrew racing game for the Nintendo Gamecube and Wii inspired by Diddy Kong Racing.

**Currently under heavy development**

# Compiling with docker

Run the following commands to build and finally run the container to compile the project.
The result will be put in the `build` folder in the root of the repository

```sh
docker build -f ./gamecube.Dockerfile -t gamecube .
docker build -f ./rehover.Dockerfile -t rehover .
docker run --volume <path to project>:/rehover --rm -it rehover
```

The final command can be re-run any time you wish to compile any changes to the project.

**Docker on windows needs special setup for [mounting](https://rominirani.com/docker-on-windows-mounting-host-directories-d96f3f056a2c)**

# Compiling from source

You will need:

- A healthy fear of the end
- devkitPro (with devkitPPC)
- Go 1.8+
- Cmake 3.1+
- [ppc-portlibs](https://github.com/Hamcha/ppc-portlibs) libraries installed, specifically:
  - entityx

## Step 0: BUILD TOOLS
Run `build.sh` or `build.cmd` (depending on your OS of choice) in the `tools` folder

## Step 1: BUILD THE PROJECT
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
