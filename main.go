package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	getText(bytes)
}
func getText(htmlBytes []byte) {

	reader := bytes.NewReader(htmlBytes)

	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			return
		}

		if tokenType == html.StartTagToken {
			tokenizer.Next()
			fmt.Println(tokenizer.Token().Data)
		}
	}
}
