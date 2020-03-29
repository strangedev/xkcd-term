package src

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func FetchPost(ID int) (*Post, error) {
	rawURL := fmt.Sprintf("https://xkcd.com/%d/", ID)
	resp, err := http.Get(rawURL)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	p := Post{ID: ID, URL: rawURL, ImageURL: "https:"}

	title := doc.Find("#ctitle").First()
	p.Title = title.Text()

	img := doc.Find("#comic img").First()
	imageURL, ok := img.Attr("src")
	if !ok {
		return nil, errors.New("can't find image URL in document")
	}
	p.ImageURL += imageURL

	imageAltText, ok := img.Attr("title")
	if !ok {
		return nil, errors.New("can't find alt text in document")
	}
	p.ImageAltText = imageAltText

	return &p, nil
}
