#!/bin/bash -eu

tools=($(sed -En 's/[[:space:]]+_ "(.*)"/\1/p' tools/cmd/cmd.go))

cd tools/cmd
echo "install go tools cmd"
for tool in ${tools[@]}
do
    echo " - Installing: $tool"
    go install "$tool"
done
