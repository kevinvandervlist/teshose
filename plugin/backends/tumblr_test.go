package backends
import (
    "testing"
    "io/ioutil"
)

func TestHtmlParser(t *testing.T) {
	tumblr := CreateTumblr("tettenvrouw")
	actual, err := tumblr.GetImageUrlFromPage(tettenvrouwPage())
	expected := "http://40.media.tumblr.com/tumblr_m3a49gP6r81rne5x6o1_500.jpg"

	if err != nil {
		t.Error("An error occurred: ", err)
	} else if actual != expected {
		t.Error("Actual image URI %s, expected %s", actual, expected)
	}
}


func tettenvrouwPage() string {
	bytes, err := ioutil.ReadFile("testdata/tettenvrouw.tumblr.com")
    if err != nil {
        return "ERROR"
    } else {
        return string(bytes)
    }
}