package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/sendhil/bookmarks/chrome"
)

func main() {
	listBookmarks := flag.Bool("list-bookmarks", true, "Whether to list the bookmarks")
	bookmark := flag.String("find-bookmark-url", "", "The Bookmark URL to find")

	flag.Parse()

	*listBookmarks = len(*bookmark) == 0

	if *listBookmarks {
		chrome.OutputBookmarks()
	} else if *bookmark != "" {
		*bookmark = strings.TrimSpace(*bookmark)
		url, err := chrome.FindURL(*bookmark)
		if err != nil {
			fmt.Println(fmt.Sprintf("Unable to find Bookmark : '%v'", *bookmark))
		} else {
			fmt.Println(url)
		}
	}
}
