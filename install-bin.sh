#!/bin/bash
if [ -z "$1" ]; then
	echo "You need to pass a version, see https://github.com/strangedev/xkcd-term/releases"
	exit 1
fi

if [ -z "$PREFIX" ]; then
    PREFIX="/usr/local"
fi


function enterWorkDir() {
	mkdir "$binary"
	pushd "$binary"
}

function cleanUpAndExit() {
	popd
	rm -rf "$binary"
	exit
}

version="$1"
binary="xkcd-linux-amd64-glibc"
checksum="xkcd-linux-amd64-glibc.sha256"
release="https://github.com/strangedev/xkcd-term/releases/download/$version/$binary"
release_checksum="$release.sha256"

enterWorkDir

curl -L "$release" -o xkcd
curl -L "$release_checksum" -O

if ! sha256sum -c "$checksum" --status; then
	echo "Checksum error!"
	cleanUpAndExit
fi

if [ $( id -u ) -eq "0" ]; then
	install xkcd "$PREFIX/bin"
else
	sudo install xkcd "$PREFIX/bin"
fi

cleanUpAndExit