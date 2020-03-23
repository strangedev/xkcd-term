# xkcd-term

A super minimalistic xkcd viewer for the terminal with support for different output formats.

I wrote this primarily to display the newest xkcds in my motd.

## Usage

```shell script
$ xkcd -help
Usage of xkcd:
  -n int
      maximum number of feed items to output (default 10)
  -o string
      controls the output format. Choose: 'human', 'json', 'yaml' (default "human")
```

**Example**: Open the latest xkcd in your default browser

```shell script
xdg-open "$( xkcd -o json | jq -r '.[0].url' )"
```
