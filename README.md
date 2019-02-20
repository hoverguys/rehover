![rehover](https://hoverguys.github.io/rehover/images/logo.svg)
<a href="https://rehover-build.ovo.ovh/docs/"><img src="https://codedocs.xyz/doxygen/doxygen.svg"/></a>

A homebrew racing game for the Nintendo Gamecube and Wii inspired by Diddy Kong Racing.

**Currently under heavy development**

# Compiling with docker

Compiling using docker will start a service container with the toolchain which can be instructed to compile the project

First start the container using the following command:

```sh
docker-compose -f "docker-compose.yml" up -d --build
```

Then instruct the container to compile:

```sh
docker exec --tty devenv make -j
```

The final command can be re-run any time you wish to compile any changes to the project.

## Visual Code

In Visual Code, a task.json file in the .vscode folder can be used to tell docker to build whenever you use the buildin build command.

```json
{
    "version": "2.0.0",
    "tasks": [
        {
            "label": "build",
            "type": "shell",
            "group": {
                "kind": "build",
                "isDefault": true
            },
            "command": "docker exec --tty devenv make -j",
            "problemMatcher": "$gcc",
            "isBackground": true,
            "presentation": {
                "panel": "dedicated",
                "showReuseMessage": false,
                "clear": true
            }
        }
    ]
}
```

The `$gcc` problemMatcher the C/C++ extension to be installed.

# Using Vagrant

First, the ugly: you need a plugin.
```
vagrant plugin install vagrant-docker-compose
```

After that, setup/run the VM with `vagrant up` and build the project with `vagrant build`.

### Why would you want to use Vagrant?

If you develop on Windows but don't want HyperV active (needed for Docker on Windows) this is a decent middle ground.

### `vagrant build` fails

Sometimes the provisioning can't keep the container running, just run `vagrant provision` and use `vagrant ssh -c "docker ps"` to check if the container is up.

# Compiling from source

You will need:

- A healthy fear of the end
- devkitPro (with devkitPPC)
- Go 1.8+
- Cmake 3.1+
- [ppc-portlibs](https://github.com/Hamcha/ppc-portlibs) libraries installed, specifically:
  - entityx

## Step 0: BUILD TOOLS
Run `tools/build.sh` or `tools/build.cmd` (depending on your OS of choice).

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
