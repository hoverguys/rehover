#!/usr/bin/env bash
cd "$(dirname "$0")/"

/rehover/tools/build.sh &&
mkdir -p /rehover/build && cd /rehover/build &&
cmake /rehover/project
make -j