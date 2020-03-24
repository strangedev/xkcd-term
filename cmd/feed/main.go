package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/strangedev/catchall"
	"gopkg.in/yaml.v2"
	"xkcd-term/src"
)

const XKCDAtom = "https://www.xkcd.com/atom.xml"

var n int
var outputFormat string
var feedURL string

func init() {
	flag.IntVar(&n, "n", 10, "maximum number of feed items to output")
	flag.StringVar(&outputFormat, "o", "human", "controls the output format. Choose: 'human', 'json', 'yaml', 'xml'")
	flag.StringVar(&feedURL, "f", XKCDAtom, "controls the feed URL in case it changes in the future")
}

func TextFormat(t string) string {
	return t
}

func URLFormat(t string) string {
	return color.New(color.FgCyan).Add(color.Underline).Sprintf(t)
}

func TitleFormat(t string) string {
	return color.New(color.Bold).Sprintf(t)
}

func main() {
	flag.Parse()

	posts, err := src.GetPosts(n, feedURL)
	catchall.CheckFatal("Can't get posts", err)

	switch outputFormat {
	case "json":
		out, err := json.Marshal(posts)
		catchall.CheckFatal("Can't encode posts", err)
		fmt.Println(string(out))
	case "yaml":
		out, err := yaml.Marshal(posts)
		catchall.CheckFatal("Can't encode posts", err)
		fmt.Println(string(out))
	case "xml":
		out, err := xml.Marshal(posts)
		catchall.CheckFatal("Can't encode posts", err)
		fmt.Println(string(out))
	case "human":
		fallthrough
	default:
		for _, post := range posts {
			fmt.Printf("%s %s \n", TitleFormat(post.Title), URLFormat(post.URL))
			fmt.Println(TextFormat(post.ImageAltText))
			fmt.Println()
		}
	}
}
