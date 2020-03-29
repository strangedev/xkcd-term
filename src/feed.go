package src

import (
	"github.com/mmcdole/gofeed"
	"golang.org/x/net/html"
	"net/url"
	"path/filepath"
	"strconv"
	"strings"
)

type Post struct {
	Title        string `json:"Title"`
	URL          string `json:"URL"`
	ImageURL     string `json:"ImageUrl"`
	ImageAltText string `json:"ImageAltText"`
	ID           int    `json:"ID"`
}

func ParseID(rawURL string) (int, error) {
	URL, err := url.Parse(rawURL)
	if err != nil {
		return 0, err
	}
	base := filepath.Base(URL.Path)
	ID, err := strconv.Atoi(base)
	if err != nil {
		return 0, err
	}
	return ID, nil
}

func GetPostsAtom(posts *[]Post, n int, feedURL string) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedURL)
	if err != nil {
		return err
	}

	for i, item := range feed.Items {
		if i >= n {
			break
		}
		ID, err := ParseID(item.GUID)
		if err != nil {
			return err
		}

		img, err := html.Parse(strings.NewReader(item.Description))
		if err != nil {
			return err
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
			ID:           ID,
			Title:        item.Title,
			ImageURL:     src,
			ImageAltText: text,
			URL:          item.GUID,
		}
		*posts = append(*posts, post)
	}
	return nil
}
