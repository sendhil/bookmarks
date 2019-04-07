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
	fileName := flag.String("json-file-name", "home.json", "The JSON File Name for Bookmarks")

	flag.Parse()

	*listBookmarks = len(*bookmark) == 0

	if *outputJSON {
		parsers.OutputChromeBookmarks(true)
	} else if *listBookmarks {
		parsers.OutputJSONBookmarks(*fileName, false)
	} else if *bookmark != "" {
		*bookmark = strings.TrimSpace(*bookmark)
		url, err := parsers.FindJSONBookmarkURL(*bookmark, *fileName)
		if err != nil {
			fmt.Println(fmt.Sprintf("Unable to find Bookmark : '%v'", *bookmark))
		} else {
			fmt.Println(url)
		}
	}
}
