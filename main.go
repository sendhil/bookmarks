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

// Node type does stuff
type Node struct {
	Key      string
	Data     interface{}
	IsFolder bool
}

func (n Node) String() string {
	return fmt.Sprintf("{Key: %v, IsFolder :%v}", n.Key, n.IsFolder)
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

	nodes := getRootNodes(roots.Roots)

	// Now parse until empty
	fmt.Println(nodes)
}

func getRootNodes(roots map[string]interface{}) []Node {
	nodes := make([]Node, 0)

	for key := range roots {
		// Is it a folder? If so parse as folder
		if isFolder(roots[key]) {
			nodes = append(nodes, Node{Key: key, Data: roots[key], IsFolder: true})
		}
	}

	return nodes
}

func isFolder(item interface{}) bool {
	if parsedItem, ok := item.(map[string]interface{}); ok {
		if val, ok := parsedItem["type"].(string); ok {
			return strings.ToLower(val) == "folder"
		}
	}

	return false
}
