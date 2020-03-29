#!/usr/bin/env bash

if [ -z "$1" ];  then
  set -e
fi

make

echo "Latest xkcd"
build/xkcd
echo
echo "Latest 5 xkcds"
build/xkcd -n 5
echo
echo "xkcd no. 420"
build/xkcd -i 420
echo
echo "7 xkcd, starting from 133"
build/xkcd -i 133 -n 7
echo
echo "3 xkcds in human format"
build/xkcd -n 3 -o human
echo
echo "3 xkcds in json format"
build/xkcd -n 3 -o json
echo
echo "3 xkcds in yaml format"
build/xkcd -n 3 -o yaml
echo
echo "3 xkcds in xml format"
build/xkcd -n 3 -o xml
echo
echo "3 xkcds, selecting ImageURL"
build/xkcd -n 3 -o select -s ImageURL
echo
echo "3 xkcds, selecting URL"
build/xkcd -n 3 -o select -s URL
echo
echo "3 xkcds, selecting Title"
build/xkcd -n 3 -o select -s Title
echo
echo "3 xkcds, selecting Caption"
build/xkcd -n 3 -o select -s Caption
echo
echo "3 xkcds, selecting ID"
build/xkcd -n 3 -o select -s ID
echo
