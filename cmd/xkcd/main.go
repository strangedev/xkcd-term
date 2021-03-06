package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"strconv"
	"time"
	"xkcd-term/src"

	"github.com/fatih/color"
	"github.com/strangedev/catchall"
	"gopkg.in/yaml.v2"
)

var n int
var outputFormat string
var feedURL string
var selectKey string
var comicID int
var cacheTimeToLive uint

func init() {
	flag.IntVar(&n, "n", 1, "maximum number of xkcds to output.")
	flag.StringVar(&outputFormat, "o", "human", "controls the output format. Choose: 'human', 'json', 'yaml', 'xml', 'select'")
	flag.StringVar(&feedURL, "f", src.XKCDAtom, "controls the atom feed URL in case it changes in the future")
	flag.StringVar(&selectKey, "s", "ImageURL", "selects value to output. For use only with 'select' output format. Choose: 'Title', 'URL', 'ImageURL', 'Caption'")
	flag.IntVar(&comicID, "i", 0, "(Optional) Selects the newest comic to output by ID. If it is 0, the atom feed is used to get the newest post.")
	flag.UintVar(&cacheTimeToLive, "t", 8, "(Optional) Number of hours after which the feed cache is marked as stale. To improve performance, the feed is only fetched every few hours. Setting -t to 0 disables the cache. This option only applies when fetching the latest comics, i.e. when -i is 0. Also, when fetching more than about 4 comics using -n, they might not all be in the cache in which case they're fetched without the cache.")
}

func textFormat(t string) string {
	return t
}

func urlFormat(t string) string {
	return color.New(color.FgCyan).Add(color.Underline).Sprintf(t)
}

func titleFormat(t string) string {
	return color.New(color.Bold).Sprintf(t)
}

func main() {
	flag.Parse()

	metas := make([]src.ComicMeta, 0, n)
	if comicID < 1 {
		err := src.GetLatestComicMetas(&metas, n, feedURL, time.Duration(cacheTimeToLive)*time.Hour)
		catchall.CheckFatal("can't read from atom feed", err)
		// we might still have metas to fetch
		n -= len(metas)
		comicID = metas[len(metas)-1].ID - 1
	}

	// this looks like it should be an else, but it might be a continuation
	// of the previous if.
	if comicID > 1 {
		for i := n; i > 0; i-- {
			meta, err := src.GetComicMeta(comicID)
			catchall.CheckFatal("can't fetch meta information", err)
			metas = append(metas, *meta)
			comicID--
		}
	}

	switch outputFormat {
	case "json":
		out, err := json.Marshal(metas)
		catchall.CheckFatal("can't encode metas as json", err)
		fmt.Println(string(out))
	case "yaml":
		out, err := yaml.Marshal(metas)
		catchall.CheckFatal("can't encode metas as yaml", err)
		fmt.Println(string(out))
	case "xml":
		out, err := xml.Marshal(metas)
		catchall.CheckFatal("can't encode metas as xml", err)
		fmt.Println(string(out))
	case "select":
		for _, meta := range metas {
			var value string
			switch selectKey {
			case "ID":
				value = strconv.Itoa(meta.ID)
			case "Title":
				value = meta.Title
			case "URL":
				value = meta.URL
			case "Caption":
				value = meta.Caption
			case "ImageURL":
				fallthrough
			default:
				value = meta.ImageURL
			}
			fmt.Println(value)
		}
	case "human":
		fallthrough
	default:
		for i, meta := range metas {
			fmt.Printf("%s %s \n", titleFormat(meta.Title), urlFormat(meta.URL))
			fmt.Println(textFormat(meta.Caption))
			if i < len(metas)-1 {
				fmt.Println()
			}
		}
	}
}
