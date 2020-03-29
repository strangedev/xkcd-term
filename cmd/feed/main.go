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
var selectKey string
var comicID int

func init() {
	flag.IntVar(&n, "n", 1, "maximum number of xkcds to output.")
	flag.StringVar(&outputFormat, "o", "human", "controls the output format. Choose: 'human', 'json', 'yaml', 'xml', 'select'")
	flag.StringVar(&feedURL, "f", XKCDAtom, "controls the feed URL in case it changes in the future")
	flag.StringVar(&selectKey, "s", "ImageURL", "selects value to output. For use only with 'select' output format. Choose: 'Title', 'URL', 'ImageURL', 'ImageAltText'")
	flag.IntVar(&comicID, "i", 0, "(Optional) Selects the newest comic to output by ID. If it is 0, the atom feed is used to get the newest post.")
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

	posts := make([]src.Post, 0, n)
	if comicID < 1 {
		err := src.GetPostsAtom(&posts, n, feedURL)
		catchall.CheckFatal("Can't read feed", err)
		// we might still have posts to fetch
		n -= len(posts)
		comicID = posts[len(posts)-1].ID - 1
	}

	if comicID > 0 {
		for i := n; i > 0; i-- {
			post, err := src.FetchPost(comicID)
			catchall.CheckFatal("Can't fetch post", err)
			posts = append(posts, *post)
			comicID--
		}
	}

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
	case "select":
		for _, post := range posts {
			var value string
			switch selectKey {
			case "Title":
				value = post.Title
			case "ImageURL":
				value = post.ImageURL
			case "ImageAltText":
				value = post.ImageAltText
			case "URL":
				fallthrough
			default:
				value = post.URL
			}
			fmt.Println(value)
		}
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
