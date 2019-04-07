package parsers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os/user"
	"strings"

	"github.com/sendhil/bookmarks/common"
)

func getJSONBookmarkFileLocation(fileName string) string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/Dropbox/Configurations/Bookmarks/%s", usr.HomeDir, fileName)
}

// OutputJSONBookmarks outputs the bookmarks stored in the user's custom folder somewhere
func OutputJSONBookmarks(fileName string, jsonOutput bool) {
	byteValue, err := ioutil.ReadFile(getJSONBookmarkFileLocation(fileName))
	if err != nil {
		panic(err)
	}

	bookmarks := []common.Bookmark{}
	err = json.Unmarshal(byteValue, &bookmarks)
	if err != nil {
		panic(err)
	}

	if jsonOutput {
		outputAsBytes, err := json.Marshal(bookmarks)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(outputAsBytes))
	} else {
		for _, bookmark := range bookmarks {
			fmt.Println(bookmark.Text)
		}
	}
}

// FindJSONBookmarkURL finds the URL based on the bookmark name
func FindJSONBookmarkURL(bookmarkText, fileName string) (string, error) {
	byteValue, err := ioutil.ReadFile(getJSONBookmarkFileLocation(fileName))
	if err != nil {
		panic(err)
	}

	bookmarks := []common.Bookmark{}
	err = json.Unmarshal(byteValue, &bookmarks)
	if err != nil {
		panic(err)
	}

	for _, bookmark := range bookmarks {
		if strings.ToLower(bookmark.Text) == strings.ToLower(bookmarkText) {
			return bookmark.URL, nil
		}
	}

	return "", errors.New("Not Found")
}
