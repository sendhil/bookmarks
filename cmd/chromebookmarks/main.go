package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/sendhil/bookmarks/parsers"
)

func main() {
	listBookmarks := flag.Bool("list-bookmarks", true, "Whether to list the bookmarks")
	bookmark := flag.String("find-bookmark-url", "", "The Bookmark URL to find")
	outputJSON := flag.Bool("output-json", false, "Whether to Output Chrome Bookmarks as JSON")

	flag.Parse()

	*listBookmarks = len(*bookmark) == 0

	if *outputJSON {
		parsers.OutputChromeBookmarks(true)
	} else if *listBookmarks {
		parsers.OutputJSONBookmarks("home.json", false)
	} else if *bookmark != "" {
		*bookmark = strings.TrimSpace(*bookmark)
		url, err := parsers.FindJSONBookmarkURL(*bookmark, "home.json")
		if err != nil {
			fmt.Println(fmt.Sprintf("Unable to find Bookmark : '%v'", *bookmark))
		} else {
			fmt.Println(url)
		}
	}
}
