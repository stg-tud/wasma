#!/bin/sh

name=$1

mkdir /usr/wasma/cmd/$name
cd /usr/wasma/cmd/$name

cp /usr/wasma/tools/scripts/template-new-analysis.txt /usr/wasma/cmd/$name/$name.go

structName=$(echo $name | tr [:lower:] [:upper:])
refName=$(echo $name | tr [:upper:] [:lower:])

sed -i 's/NEWANALYSIS/'$structName/ /usr/wasma/cmd/$name/$name.go
sed -i 's/nEWANALYSIS/'$refName/ /usr/wasma/cmd/$name/$name.go

nano /usr/wasma/cmd/$name/$name.go