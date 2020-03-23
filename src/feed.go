package src

import (
	"github.com/mmcdole/gofeed"
	"golang.org/x/net/html"
	"strings"
)

const XKCDAtom = "https://www.xkcd.com/atom.xml"
const FeedSize = 4

type Post struct {
	Title string `json:"title"`
	ImageURL string `json:"imageURL"`
	Text string `json:"text"`
}

func GetPosts(n int) ([]Post, error) {
	posts := make([]Post, 0, FeedSize)
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(XKCDAtom)
	if err != nil {
		return nil, err
	}

	for i, item := range feed.Items {
		if i >= n {
			break
		}
		img, err := html.Parse(strings.NewReader(item.Description))
		if err != nil {
			return nil, err
		}

		var text, src string
		var crawler func(*html.Node)
		crawler = func(node *html.Node) {
			if node.Type == html.ElementNode && node.Data == "img" {
				for _, attr := range node.Attr {
					if attr.Key == "title" {
						text = attr.Val
						continue
					}
					if attr.Key == "src" {
						src = attr.Val
						continue
					}
				}
				return
			}
			for child := node.FirstChild; child != nil; child = child.NextSibling {
				crawler(child)
			}
		}
		crawler(img)

		post := Post{
			Title:    item.Title,
			ImageURL: src,
			Text:     text,
		}
		posts = append(posts, post)
	}
	return posts, nil
}
