# xkcd-term

![Build Status](https://github.com/strangedev/xkcd-term/workflows/CI/badge.svg)

A super minimalistic xkcd viewer for the terminal with support for different output formats.

I wrote this primarily to display the newest xkcds in my motd.

![A screenshot of the human readable output](https://i.postimg.cc/Hnvbx4Sm/2020-03-24-005001-762x341-scrot.png)

## Usage

```shell script
$ xkcd -help
Usage of build/xkcd:
  -f string
    	controls the feed URL in case it changes in the future (default "https://www.xkcd.com/atom.xml")
  -n int
    	maximum number of feed items to output (default 10)
  -o string
    	controls the output format. Choose: 'human', 'json', 'yaml', 'xml' (default "human")
  -s string
    	selects key to output. For use only with 'select' output format. Choose: 'Title', 'URL', 'ImageURL', 'ImageAltText' (default "ImageURL")
```

**Example**: Open the latest xkcd in your default browser

```shell script
xdg-open "$( xkcd -o json | jq -r '.[0].url' )"
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