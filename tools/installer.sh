#!/bin/bash -eu

tools=($(sed -En 's/[[:space:]]+_ "(.*)"/\1/p' tools/tools.go))

cd tools
echo "install go tools"
for tool in ${tools[@]}
do
    echo " - Installing: $tool"
    go install "$tool"
done
