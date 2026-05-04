package main

import (
	"bufio"
	"os"
	"encoding/hex"
	"fmt"
	"strings"
	"encoding/json"
	"crypto/md5"
)

const urlsFilePath = "urlsrepo/urls.json"

type urlsRepo map[string]string

func main() {
	urlsFile, err := os.OpenFile(urlsFilePath, os.O_RDWR|os.O_CREATE, 0644)
	defer urlsFile.Close()

	if err != nil {
		fmt.Println("Error opening urls file:", err)
		return
	}

	// Create a reader from the byte slice
	decoder := json.NewDecoder(urlsFile)
	var savedUrls urlsRepo
	if err := decoder.Decode(&savedUrls); err != nil {
		fmt.Println("Error unmarshalling savedurls:", err)
		return
	}

	for {
		fmt.Println("Want a url to be shortened or want the long url:")
		
		buf := bufio.NewReader(os.Stdin)

		command, err := buf.ReadString('\n')
		
		if err != nil {
			fmt.Println("Error reading command:", err)
			return
		}

		command = strings.TrimSpace(command)

		commandWords := strings.Fields(command)

		if commandWords[0] == "quit" {
			fmt.Println("Quitting...")
			return
		}

		if commandWords[0] == "shorten" {
			url := commandWords[1]

			hash := md5.Sum([]byte(url))
			savedUrls[hex.EncodeToString(hash[:])] = url
			
			// temp := string(hash[:])
			// fmt.Println("Saved url: ", temp)
			// fmt.Printf("%x", hash)

			// use Using json.Encoder to write to the file now
			encoder := json.NewEncoder(urlsFile)
 			if err := encoder.Encode(savedUrls); err != nil {
				fmt.Println("Error writing saved Urls file:", err)
				return
			}
		} else if commandWords[0] == "get" {
			hash := commandWords[1]

			fullUrl, exists := savedUrls[hash]
			if !exists {
				fmt.Println("invalid key: URL not found")
				continue
			}
			fmt.Println("The long url is: ", fullUrl)
		} else {
			fmt.Println("Unknown command:", commandWords[0])
		}
	}
}
