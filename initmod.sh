#!/usr/bin/env bash

pwd=`pwd`
for x in */ ; do
  cd $x
  go mod init github.com/kokizzu/gotro/${x%?} # remove lasat character
  cd $pwd
done 

for x in */ ; do
  go mod tidy
done
