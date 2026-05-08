package main

import (
	"bufio"
	"strconv"
	"fmt"
	"os"
	"strings"
	"lru-cache/dell"
)

func main() {
	// let's make it a console app
	// before that, let's test this
	// linked list first
	

	reader := bufio.NewReader(os.Stdin)

	var size int

	
	fmt.Println("Enter the size of the cache:")
	sizeStr, _ := reader.ReadString('\n')
	size, _ = strconv.Atoi (strings.TrimSpace(sizeStr))

	fmt.Println("Size:", size)
	
	dq := dell.New(size)

	hashMap := make(map[string]*dell.Node)

	for {
		fmt.Println("Enter a command out of get, put:")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)
		if (command != "get" && command != "put" && command != "print") {
			fmt.Println("Invalid command")
			continue
		}

		if (command == "get") {
			fmt.Println("Enter the key:")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)
			node, exists := hashMap[key]
			if (exists) {
				dq.MoveNodeToLast(node)
				fmt.Println("Here is the value:", node.Value)
			} else {
				fmt.Println("Key not found")
			}
			continue
		}

		if (command == "print") {
			dq.Print()
			continue
		}

		if (command == "put") {
			fmt.Println("Enter the Key & Value in separate lines:")
			key, _ := reader.ReadString('\n')
			value, _ := reader.ReadString('\n')
			valueInt, _ := strconv.Atoi(strings.TrimSpace(value))
			key = strings.TrimSpace(key)

			node, exists := hashMap[key]
			if (exists) {
				// refresh it's value
				node.Value = valueInt
				dq.MoveNodeToLast(node)
				fmt.Println("Value refreshed successfully")
				continue
			}
			// know if delete happened or not
			// if yes then which key was deleted and delete that from here as well
			if (dq.Length >= size) {
				// need to know which key will be deleted here
				firstNode, _ := dq.GetFirstNode()
				delete(hashMap, firstNode.Key)
			}
			node, err := dq.Add(key, valueInt)
			if (err != nil) {
				fmt.Println("Error adding value:", err)
				continue
			}
			hashMap[key] = node
			fmt.Println("Value added successfully")
		}

	}
	
}