# xkcd-term

![Build Status](https://github.com/strangedev/xkcd-term/workflows/CI/badge.svg)

A super minimalistic xkcd viewer for the terminal with support for different output formats.

I wrote this primarily to display the newest xkcds in my motd.

![A screenshot of the human readable output](https://i.postimg.cc/Hnvbx4Sm/2020-03-24-005001-762x341-scrot.png)

## Usage

```
$ xkcd -help
Usage of build/xkcd:
  -f string
    	controls the feed URL in case it changes in the future (default "https://www.xkcd.com/atom.xml")
  -i int
    	(Optional) Selects the newest comic to output by ID. If it is 0, the atom feed is used to get the newest post.
  -n int
    	maximum number of xkcds to output. (default 1)
  -o string
    	controls the output format. Choose: 'human', 'json', 'yaml', 'xml', 'select' (default "human")
  -s string
    	selects value to output. For use only with 'select' output format. Choose: 'Title', 'URL', 'ImageURL', 'ImageAltText' (default "ImageURL")
```

#### Examples

Open the latest xkcd in your default browser

```shell script
xdg-open $( xkcd -o select -s URL )
```

View the last 5 xkcds with `feh`

```shell script
xkcd -n 5 -o select -s ImageURL | xargs -n 1 feh
```

View xkcd number 2000

```shell script
xkcd -i 2000
```

View xkcds 190 - 200:

```shell script
xkcd -i 200 -n 11
```

Display the latest xkcd caption and URL with cowsay

```shell script
#!/bin/bash
xkcd_latest=$( xkcd -o json )

echo "$xkcd_latest" | jq '.[].Caption' | xargs cowsay -f three-eyes
echo "$xkcd_latest" | jq '.[].ImageURL'
```


## Install

Get your [hot, fresh, and very chewy pre-built binaries for GNU/Linux](https://github.com/strangedev/xkcd-term/releases),
or build from source. Go `v1.14` is recommended for building.

```shell script
git clone https://github.com/strangedev/xkcd-term.git
cd xkcd-term
make
sudo make install
```

On Linux, you can use the `install-bin.sh` bash script to download
and verify releases:

```shell script
./install-bin.sh v2.1.0
```
