package src

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/mmcdole/gofeed"
	"github.com/strangedev/catchall"
	"golang.org/x/net/html"
)

const (
	CacheDir      = ".xkcd"
	FeedFile      = "feed.xml"
	CacheInfoFile = "info.json"
	XKCDAtom      = "https://www.xkcd.com/atom.xml"
)

type CacheInfo struct {
	ModifiedAt time.Time `json:modified_at`
}

func getCachePath() string {
	home, err := homedir.Dir()
	catchall.CheckFatal("Can't get the user home directory.", err)

	return filepath.Join(home, CacheDir)
}

func ensureCacheDirectoryExists() error {
	if _, err := os.Stat(getCachePath()); os.IsNotExist(err) {
		if err = os.Mkdir(getCachePath(), 0755); err != nil {
			return err
		}
	}

	return nil
}

func readCacheInfo() (info CacheInfo, err error) {
	infoPath := filepath.Join(getCachePath(), CacheInfoFile)
	infoFile, err := os.Open(infoPath)
	if err != nil {
		if os.IsNotExist(err) {
			return CacheInfo{ModifiedAt: time.Unix(0, 0)}, nil
		}
		return
	}
	defer infoFile.Close()

	bytes, err := ioutil.ReadAll(infoFile)
	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &info)
	return
}

func updateFeed(feedURL string, cacheTimeToLive time.Duration) error {
	if err := ensureCacheDirectoryExists(); err != nil {
		return err
	}

	cacheInfo, err := readCacheInfo()
	if err == nil {
		cacheAge := time.Now().Sub(cacheInfo.ModifiedAt)
		if cacheAge < cacheTimeToLive {
			return nil
		}
	}

	feedPath := filepath.Join(getCachePath(), FeedFile)
	resp, err := http.Get(feedURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	feedFile, err := os.Create(feedPath)
	if err != nil {
		return err
	}
	defer feedFile.Close()

	infoPath := filepath.Join(getCachePath(), CacheInfoFile)
	infoFile, err := os.Create(infoPath)
	if err != nil {
		return err
	}
	defer infoFile.Close()

	_, err = io.Copy(feedFile, resp.Body)
	if err != nil {
		return err
	}
	cacheInfo = CacheInfo{ModifiedAt: time.Now()}
	infoJSON, err := json.Marshal(cacheInfo)
	if err != nil {
		return err
	}
	infoFile.Write(infoJSON)

	return err
}

func GetLatestComicMetas(metas *[]ComicMeta, nMax int, feedURL string, cacheTimeToLive time.Duration) error {
	err := updateFeed(feedURL, cacheTimeToLive)
	if err != nil {
		return err
	}

	feedPath := filepath.Join(getCachePath(), FeedFile)
	feedBytes, err := ioutil.ReadFile(feedPath)
	if err != nil {
		return err
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(feedBytes))
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
		src, ok = img.Attr("src")
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
