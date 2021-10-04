#!/bin/sh

cd /usr/wasma/cmd

for file in $(find $(pwd) -type f -name "*.go");
do
  file=${file##*/} # remove path
  file=${file%.go} # remove extension

  echo "build -> $file"
  go build -o /usr/wasma/bin/$file /usr/wasma/cmd/$file/$file.go
done