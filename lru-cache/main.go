package main

import (
	"bufio"
	"strconv"
	"fmt"
	"os"
	"strings"
	"lru-cache/cache"
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
	
	c := cache.New(size)

	for {
		fmt.Println("Enter a command out of get, put, print:")
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
			value, err := c.Get(key)
			if (err != nil) {
				fmt.Println("Error getting value:", err)
				continue
			}
			fmt.Println("Here is the value:", value)
			continue
		}

		if (command == "print") {
			c.Print()
			continue
		}

		if (command == "put") {
			fmt.Println("Enter the Key & Value in separate lines:")
			key, _ := reader.ReadString('\n')
			value, _ := reader.ReadString('\n')
			valueInt, _ := strconv.Atoi(strings.TrimSpace(value))
			key = strings.TrimSpace(key)

			err := c.Put(key, valueInt)
			if (err != nil) {
				fmt.Println("Error adding value:", err)
				continue
			}
			fmt.Println("Value added successfully")
		}

	}
	
}