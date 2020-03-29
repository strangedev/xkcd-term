package src

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// GetComicMeta fetches an xkcd by ID, parses the page, and returns the comic's meta information
func GetComicMeta(ID int) (*ComicMeta, error) {
	rawURL := fmt.Sprintf("https://xkcd.com/%d/", ID)
	resp, err := http.Get(rawURL)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	p := ComicMeta{ID: ID, URL: rawURL, ImageURL: "https:"}

	title := doc.Find("#ctitle")
	if len(title.Nodes) != 1 {
		if len(title.Nodes) < 1 {
			return nil, errors.New("can't find title in document")
		} else {
			return nil, errors.New(fmt.Sprintf("ambiguous title, got %d nodes", len(title.Nodes)))
		}
	}
	p.Title = title.Text()

	img := doc.Find("#comic img").First()
	imageURL, ok := img.Attr("src")
	if !ok {
		return nil, errors.New("can't find image URL in document")
	}
	p.ImageURL += imageURL

	caption, ok := img.Attr("title")
	if !ok {
		return nil, errors.New("can't find caption in document")
	}
	p.Caption = caption

	return &p, nil
}
