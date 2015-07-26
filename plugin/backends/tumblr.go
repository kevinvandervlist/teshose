package backends
import (
	"io/ioutil"
	"net/http"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"os"
	"path"
)

type Tumblr struct {
	title string
	tempFolder string
}

func CreateTumblr(tumblr string) (*Tumblr) {
	return &Tumblr{
		title: tumblr,
		tempFolder: os.TempDir(),
	}
}

func (tumblr *Tumblr) GetName() string {
	return tumblr.title
}

func (tumblr *Tumblr) GetRandomPage() (string, error) {
	source := "http://tettenvrouw.tumblr.com/random"
	response, err := http.Get(source)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	page, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}
	return string(page), nil
}

func (tumblr *Tumblr) GetImageUrlFromPage(page string) (string, error) {
	switch tumblr.title {
	case "tettenvrouw":
		return tettenvrouw(page)
	default:
		return "", errors.New("No implementation found for tumblr " + tumblr.title)
	}
}

func (tumblr *Tumblr) DownloadImage(url string) (string, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	filename := path.Base(url)

	//fh, err := ioutil.WriteFile(tumblr.tempFolder + os.PathListSeparator + tumblr.title + , contents, 0644)
	//fh, err := ioutil.TempFile(tumblr.tempFolder, tumblr.title)
	//defer fh.Close()

	if err != nil {
		return "", err
	}
	//path := fh.Name()
	//_, err = fh.Write(contents)
	path := tumblr.tempFolder + "/" + tumblr.title + filename
	err = ioutil.WriteFile(path, contents, 0644)
	if err != nil {
		return "", err
	} else {
		return path, nil
	}
}

func tettenvrouw(page string) (string, error) {
	reader := strings.NewReader(page)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", err
	}
	selection := doc.Find(".main_photo").EachWithBreak(func (i int, s *goquery.Selection) (bool) {
		_, exists := s.Attr("src")
		return exists
	})
	uri, exists := selection.Attr("src")
	if exists {
		return uri, nil
	} else {
		return "", errors.New("No URI found.")
	}
}