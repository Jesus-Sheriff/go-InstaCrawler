package main

import (
	"log"
	"os"

	"gopkg.in/ahmdrz/goinsta.v2"
)

func main() {
	insta := goinsta.New(
		os.Getenv("INSTAGRAM_USERNAME"),
		os.Getenv("INSTAGRAM_PASSWORD"),
	)

	if err := insta.Login(); err != nil {
		log.Println(err)
		return
	}
	defer insta.Logout()

	feedTag, err := insta.Feed.Tags("golang")
	if err != nil {
		log.Println(err)
		return
	}

	//log.Printf("ultima foto: ", feedTag.Images[0].User.Username)
	log.Printf("ultima foto:  %s", feedTag.Images[0].Images.GetBest())
	// for _, item := range feedTag.RankedItems {
	// 	user := item.User
	// 	user.SetInstagram(insta)

	// 	err = user.Follow()
	// 	if err != nil {
	// 		log.Printf("error on following user %s, %v", user.Username, err)
	// 	} else {
	// 		log.Printf("start following user %s", user.Username)
	// 	}
	// }
}