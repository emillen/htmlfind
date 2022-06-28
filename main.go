package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ericchiang/css"
	"golang.org/x/net/html"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)

	if len(os.Args) < 3 {
		panic(errors.New("not enough arguments"))
	}

	if err != nil {
		panic(err)
	}

	if os.Args[1] == "text" {
		getText(bytes, strings.Join(os.Args[2:], " "))
	}

	if os.Args[1] == "attr" {

		getAttribs(bytes, os.Args[2:])
	}
}

func getAttribs(htmlBytes []byte, attributes []string) {
	reader := bytes.NewReader(htmlBytes)
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {

			return
		}

		if tokenType == html.StartTagToken || tokenType == html.SelfClosingTagToken {
			for _, attribute := range tokenizer.Token().Attr {
				if contains(attributes, attribute.Key) {
					fmt.Println(attribute.Val)
				}
			}

		}

	}
}

func getText(htmlBytes []byte, selector string) {
	sel, err := css.Parse(selector)
	if err != nil {
		panic(err)
	}
	reader := bytes.NewReader(htmlBytes)
	node, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}
	for _, ele := range sel.Select(node) {
		fmt.Println(ele.FirstChild.Data)
	}
	fmt.Println()
}

func contains(stringSlice []string, s string) bool {
	for _, item := range stringSlice {
		if item == s {
			return true
		}
	}
	return false
}
