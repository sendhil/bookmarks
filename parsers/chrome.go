package parsers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os/user"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/sendhil/bookmarks/common"
)

func getChromeBookmarksFolder() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/.config/google-chrome/Default/Bookmarks", usr.HomeDir)
}

// FindURL finds the URL based on the bookmark name
func FindChromeBookmarkURL(bookmark string) (string, error) {
	byteValue, err := ioutil.ReadFile(getChromeBookmarksFolder())
	if err != nil {
		panic(err)
	}

	nodes := make([][]byte, 0)

	// Initial seed
	err = jsonparser.ObjectEach(byteValue, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if dataType == jsonparser.Object {
			nodes = append(nodes, value)
		}
		return nil
	}, "roots")

	if err != nil {
		panic(err)
	}

	for len(nodes) > 0 {
		if nodeType, err := jsonparser.GetString(nodes[0], "type"); err == nil {
			keyVal, err := jsonparser.GetString(nodes[0], "name")
			if err != nil {
				fmt.Println("Error parsing:", nodes[0])
				break
			}

			if nodeType == "folder" {
				_, err = jsonparser.ArrayEach(nodes[0], func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
					nodes = append(nodes, value)
				}, "children")

				if err != nil {
					fmt.Println("Error parsing children for:", string(nodes[0]))
				}
			} else if nodeType == "url" {
				if strings.ToLower(keyVal) == strings.ToLower(bookmark) {
					strVal, err := jsonparser.GetString(nodes[0], "url")
					if err != nil {
						fmt.Println("Error parsing node : ", string(nodes[0]))
						return "", err
					}

					return strVal, nil
				}
			} else {
				fmt.Println("Unknown node type: ", nodeType)
			}
		} else {
			fmt.Println("Unable to parse node : ", string(nodes[0]))
			fmt.Println("err:", err)
		}

		nodes = append(nodes[1:])
	}

	return "", errors.New("Not Found")
}

// OutputChromeBookmarks outputs the bookmarks stored by Chrome
func OutputChromeBookmarks(jsonOutput bool) {
	byteValue, err := ioutil.ReadFile(getChromeBookmarksFolder())
	if err != nil {
		panic(err)
	}

	nodes := make([][]byte, 0)

	// Initial seed
	err = jsonparser.ObjectEach(byteValue, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if dataType == jsonparser.Object {
			nodes = append(nodes, value)
		}
		return nil
	}, "roots")

	if err != nil {
		panic(err)
	}

	bookmarks := make([]common.Bookmark, 0)

	for len(nodes) > 0 {
		if nodeType, err := jsonparser.GetString(nodes[0], "type"); err == nil {
			keyVal, err := jsonparser.GetString(nodes[0], "name")
			if err != nil {
				fmt.Println("Error parsing:", nodes[0])
				break
			}

			if nodeType == "folder" {
				_, err = jsonparser.ArrayEach(nodes[0], func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
					nodes = append(nodes, value)
				}, "children")

				if err != nil {
					fmt.Println("Error parsing children for:", string(nodes[0]))
				}
			} else if nodeType == "url" {
				urlVal, err := jsonparser.GetString(nodes[0], "url")
				if err != nil {
					fmt.Println("Error parsing:", nodes[0])
					break
				}
				bookmarks = append(bookmarks, common.Bookmark{Text: keyVal, URL: urlVal})
			} else {
				fmt.Println("Unknown node type: ", nodeType)
			}
		} else {
			fmt.Println("Unable to parse node : ", string(nodes[0]))
			fmt.Println("err:", err)
		}

		nodes = append(nodes[1:])
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
