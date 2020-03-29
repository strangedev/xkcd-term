package src

import (
	"io"
	"net/url"
	"path/filepath"
	"strconv"
)

// ComicMeta is the meta information associated with an xkcd
type ComicMeta struct {
	Title    string `json:"Title"`
	URL      string `json:"URL"`
	ImageURL string `json:"ImageURL"`
	Caption  string `json:"Caption"`
	ID       int    `json:"ID"`
}

type Comic interface {
	// implements io.ReaderCloser to read image from
	io.ReadCloser
	// every comic has some meta-information associated with it, eg. title, number, caption
	Meta() *ComicMeta
}

// ParseID returns the ID contained in an xkcd url
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

