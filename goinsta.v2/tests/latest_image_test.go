package tests

import (
	"os"
	"testing"
	"net/url"

	"gopkg.in/ahmdrz/goinsta.v2"
)

func TestImportAccount(t *testing.T) {
	insta := goinsta.New(
		os.Getenv("INSTAGRAM_USERNAME"),
		os.Getenv("INSTAGRAM_PASSWORD"),
	)
	if err := insta.Login(); err != nil {
		t.Fatal(err)
		return
	}
	defer insta.Logout()

	feedTag, err := insta.Feed.Tags("golang")
	if err != nil {
		t.Fatal(err)
		return
	}
	// URL de ejemplo:
	// https://scontent-mad1-1.cdninstagram.com/vp/7c8004a33e8ef83675e7c62a62c821d7/5E39388E/t51.2885-15/e35/70513351_167187977761265_1918517610523590583_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=105&se=8&ig_cache_key=MjE1MDc2MzEwMzMzMTk0ODE0Mw%3D%3D.2

	//chequeo de que la URL devuelta es válida en estructura
	bestImageURL := feedTag.Images[0].Images.GetBest()
	_, err = url.ParseRequestURI(bestImageURL)
    if err != nil {
        t.Fatal(err)
		return
    }

	t.Logf("URL is: %s", bestImageURL)

	t.Logf("logged into Instagram as user '%s'", insta.Account.Username)
}
