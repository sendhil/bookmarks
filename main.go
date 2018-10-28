package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

// Root type is the main type
type Root struct {
	Roots map[string]interface{} `json:"roots"`
}

func main() {
	fileName := "/home/sendhil/.config/google-chrome/Default/Bookmarks"
	byteValue, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var roots Root

	err = json.Unmarshal(byteValue, &roots)
	if err != nil {
		panic(err)
	}

	for key := range roots.Roots {
		if isFolder(roots.Roots[key]) {
			fmt.Println(key)
		}
	}
}

func isFolder(item interface{}) bool {
	if parsedItem, ok := item.(map[string]interface{}); ok {
		if val, ok := parsedItem["type"].(string); ok {
			return strings.ToLower(val) == "folder"
		}
	}

	return false
}
