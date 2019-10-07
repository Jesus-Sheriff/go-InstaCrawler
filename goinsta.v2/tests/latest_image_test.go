package tests

import (
	"os"
	"testing"

	"github.com/ahmdrz/goinsta/v2"
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

	// 	feedTag.Images[0].Images.GetBest()
	// //Comprobar imagenes

	t.Logf("logged into Instagram as user '%s'", insta.Account.Username)
}
