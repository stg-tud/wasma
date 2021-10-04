#!/bin/sh

# Converts all dot files in a directory to images.

pathDotFiles=$1 # Path were the dot files are stored.
imgFormat=$2 # Image format, for example, png.

# Get a list of all dot files.
dotFiles=$(find $pathDotFiles -type f -name "*.dot")

# Loop over all dot files
for dotFile in $dotFiles
do
  dot -T$imgFormat $dotFile -o ${dotFile%.dot}.$imgFormat
done
