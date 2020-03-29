package src

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
	"golang.org/x/net/html"
	"strings"
)

const XKCDAtom = "https://www.xkcd.com/atom.xml"

func GetLatestComicMetas(metas *[]ComicMeta, nMax int, feedURL string) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedURL)
	if err != nil {
		return err
	}

	for i, item := range feed.Items {
		if i >= nMax {
			break
		}
		ID, err := ParseID(item.GUID)
		if err != nil {
			return err
		}

		root, err := html.Parse(strings.NewReader(item.Description))
		if err != nil {
			return err
		}

		var title, src string

		doc := goquery.NewDocumentFromNode(root)
		img := doc.Find("img").First()
		title, ok := img.Attr("title")
		if !ok {
			return errors.New("can't find caption")
		}
		src, ok = img.Attr("title")
		if !ok {
			return errors.New("can't find image url")
		}

		post := ComicMeta{
			ID:       ID,
			Title:    item.Title,
			ImageURL: src,
			Caption:  title,
			URL:      item.GUID,
		}
		*metas = append(*metas, post)
	}
	return nil
}
