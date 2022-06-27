package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)

	if len(os.Args) < 3 {
		log.Fatal(errors.New("not enough arguments"))
	}
	if err != nil {
		log.Fatal(err)
	}
	if os.Args[1] == "text" {
		getText(bytes, os.Args[2:])
	}
}

func getText(htmlBytes []byte, tagNames []string) {

	reader := bytes.NewReader(htmlBytes)

	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			return
		}

		if tokenType == html.StartTagToken {
			if contains(tagNames, tokenizer.Token().Data) {
				tokenizer.Next()
				fmt.Println(tokenizer.Token().Data)
			}
		}
	}
}

func contains(stringSlice []string, s string) bool {
	for _, item := range stringSlice {
		if item == s {
			return true
		}
	}
	return false
}
