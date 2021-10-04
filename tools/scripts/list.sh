#!/bin/sh

parameter=$1

if [ "$parameter" = "-b" ];
then
  ls /usr/wasma/bin
else
  ls /usr/wasma/cmd
fi