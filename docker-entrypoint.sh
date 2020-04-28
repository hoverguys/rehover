#!/usr/bin/env bash
cd "$(dirname "$0")/"

cd /workspaces/rehover
./tools/build.sh &&
mkdir -p ./build && cd ./build &&
cmake ../project
make -j
