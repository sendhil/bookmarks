package main

import (
	"fmt"
	"io/ioutil"

	"github.com/buger/jsonparser"
)

func main() {
	fileName := "/home/sendhil/.config/google-chrome/Default/Bookmarks"
	byteValue, err := ioutil.ReadFile(fileName)
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

	bookmarks := make([]string, 0)

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
				bookmarks = append(bookmarks, keyVal)
			} else {
				fmt.Println("Unknown node type: ", nodeType)
			}
		} else {
			fmt.Println("Unable to parse node : ", string(nodes[0]))
			fmt.Println("err:", err)
		}

		nodes = append(nodes[1:])
	}

	for _, bookmark := range bookmarks {
		fmt.Println(bookmark)
	}
}
