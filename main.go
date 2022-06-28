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

const help = `
Takes html on stdin, find things in it, and print it to stdout

Usage: htmlfind <command> [<args>]

Commands
	text <css-selector>	Return the text in the elements returned from the css-selector
	attr [<attribute>]	Return the value of the attributes
	comments		Return all HTML comments found in the documents
`

func main() {
	file := os.Stdin
	fi, _ := file.Stat()

	if fi.Size() <= 0 {
		fmt.Print(help)
		return
	}

	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(errors.New("Failed to read from stdin"))
	}

	if len(os.Args) >= 3 && os.Args[1] == "text" {
		getText(bytes, strings.Join(os.Args[2:], " "))
		return
	}

	if len(os.Args) >= 3 && os.Args[1] == "attr" {

		getAttribs(bytes, os.Args[2:])
		return
	}

	if len(os.Args) >= 2 && os.Args[1] == "comments" {

		getComments(bytes)
		return
	}

	fmt.Print(help)
}

func getComments(htmlBytes []byte) {
	reader := bytes.NewReader(htmlBytes)
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {

			return
		}

		if tokenType == html.CommentToken {
			fmt.Println(tokenizer.Token().Data)
		}

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
		if ele.FirstChild != nil {

			fmt.Println(ele.FirstChild.Data)
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
