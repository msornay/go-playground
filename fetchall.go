// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	file, err := os.Open("data/top-1m.csv")
	if err != nil {
		fmt.Println("Error: %v", err)
		return
	}

	reader := csv.NewReader(file)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		url := fmt.Sprintf("http://www.%s", line[1])
		go fetch(url)
	}

	file.Close()
}

func fetch(url string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("Error %s: %v\n", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("%s %7d %.2fs\n", url, nbytes, secs)
}

//!-
