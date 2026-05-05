package main

import (
	// "bufio"
	// "os"
	// "fmt"
	// "io/fs"
	// "sync"
	"strings"
	"strconv"
	// "sync/atomic"
	"net/http"
	// "encoding/json"
)

func main() {

	numericRomainMapping := map[int]string{
		1: "I",
		2: "II",
		3: "III",	
		4: "IV",
		5: "V",
		6: "VI",
		7: "VII",
		8: "VIII",
		9: "IX",
		10: "X",
	}

	http.HandleFunc("/convert/", func(w http.ResponseWriter, r *http.Request) {

		numberToConvert := strings.Split(r.URL.Path, "/")
		numberToConvertInt, err := strconv.Atoi(numberToConvert[len(numberToConvert)-1])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid number"))
			return
		}

		w.Write([]byte("Number to convert: " + numericRomainMapping[numberToConvertInt]))

		w.WriteHeader(http.StatusOK)

	})

	s := http.Server {
		Addr: ":8080",
	}

	s.ListenAndServe()
	

}
