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
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {

		fmt.Println("No input on stdin")
		fmt.Print(help)
		os.Exit(1)
	}

	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(errors.New("Failed to read from stdin"))
	}

	if len(os.Args) >= 3 && os.Args[1] == "text" {
		printText(bytes, strings.Join(os.Args[2:], " "))
		return
	}

	if len(os.Args) >= 3 && os.Args[1] == "attr" {

		printAttribs(bytes, os.Args[2:])
		return
	}

	if len(os.Args) >= 2 && os.Args[1] == "comments" {

		printComments(bytes)
		return
	}

	fmt.Print(help)
	os.Exit(2)
}

func printComments(htmlBytes []byte) {
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

func printAttribs(htmlBytes []byte, attributes []string) {
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
func traverseAndPrint(node *html.Node) {
	if node == nil {
		return
	}
	if node.Type == html.TextNode {
		fmt.Print(node.Data)
	}
	if node.Type == html.ElementNode {
		fmt.Println()
		traverseAndPrint(node.FirstChild)
	}
	traverseAndPrint(node.NextSibling)
}

func printText(htmlBytes []byte, selector string) {
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

			traverseAndPrint(ele.FirstChild)
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
