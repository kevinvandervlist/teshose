package backends
import (
    "testing"
    "io/ioutil"
)

func TestHtmlParserTettenvrouw(t *testing.T) {
	tumblr := CreateTumblr("tettenvrouw")
	actual, err := tumblr.GetImageUrlFromPage(localPage("tettenvrouw"))
	expected := "http://40.media.tumblr.com/tumblr_m3a49gP6r81rne5x6o1_500.jpg"

	if err != nil {
		t.Error("An error occurred: ", err)
	} else if actual != expected {
		t.Error("Actual image URI %s, expected %s", actual, expected)
	}
}

func TestHtmlParserLingeriebomb(t *testing.T) {
	tumblr := CreateTumblr("lingeriebomb")
	actual, err := tumblr.GetImageUrlFromPage(localPage("lingeriebomb"))
	expected := "http://41.media.tumblr.com/tumblr_m7zixwjGOW1r8bjnao1_1280.jpg"

	if err != nil {
		t.Error("An error occurred: ", err)
	} else if actual != expected {
		t.Error("Actual image URI %s, expected %s", actual, expected)
	}
}


func localPage(page string) string {
	bytes, err := ioutil.ReadFile("testdata/" + page + ".tumblr.com")
	if err != nil {
		return "ERROR"
	} else {
		return string(bytes)
	}
}