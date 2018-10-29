package main

import (
	"fmt"
	"io/ioutil"

	"github.com/buger/jsonparser"
)

// Node type does stuff
type Node struct {
	Key  string
	Data []byte
}

func (n Node) String() string {
	return fmt.Sprintf("{Key: %v}", n.Key)
}

func main() {
	fileName := "/home/sendhil/.config/google-chrome/Default/Bookmarks"
	byteValue, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	nodes := make([]Node, 0)

	// Initial seed
	err = jsonparser.ObjectEach(byteValue, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
		if dataType == jsonparser.Object {
			nodes = append(nodes, Node{Key: string(key), Data: value})
		}
		return nil
	}, "roots")

	if err != nil {
		panic(err)
	}

	bookmarks := make([]string, 0)

	for len(nodes) > 0 {
		if nodeType, err := jsonparser.GetString(nodes[0].Data, "type"); err == nil {
			keyVal, err := jsonparser.GetString(nodes[0].Data, "name")
			if err != nil {
				fmt.Println("Error parsing:", nodes[0].Data)
				break
			}

			if nodeType == "folder" {
				_, err = jsonparser.ArrayEach(nodes[0].Data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
					nodes = append(nodes, Node{Data: value})
				}, "children")

				if err != nil {
					fmt.Println("Error parsing children for:", string(nodes[0].Data))
				}
			} else if nodeType == "url" {
				bookmarks = append(bookmarks, keyVal)
			} else {
				fmt.Println("Unknown node type: ", nodeType)
			}
		} else {
			fmt.Println("Unable to parse node : ", string(nodes[0].Data))
			fmt.Println("err:", err)
		}

		nodes = append(nodes[1:])
	}

	for _, bookmark := range bookmarks {
		fmt.Println(bookmark)
	}
}

// func getNodeData(node Node, isFolder bool) Node {
// 	nodeData := node.Data.(map[string]interface{})
// 	var children interface{}
// 	if childData, ok := nodeData["children"]; ok {
// 		children = childData
// 	}
//
// 	return Node{Key: nodeData["name"].(string), Data: children, IsFolder: isFolder}
// }
//
// func getRootNodes(roots map[string]interface{}) []Node {
// 	nodes := make([]Node, 0)
//
// 	for key := range roots {
// 		// Is it a folder? If so parse as folder
// 		if isFolder(roots[key]) {
// 			nodes = append(nodes, Node{Key: key, Data: roots[key], IsFolder: true})
// 		}
// 	}
//
// 	return nodes
// }
//
// func isFolder(item interface{}) bool {
// 	if parsedItem, ok := item.(map[string]interface{}); ok {
// 		if val, ok := parsedItem["type"].(string); ok {
// 			return strings.ToLower(val) == "folder"
// 		}
// 	}
//
// 	return false
// }
