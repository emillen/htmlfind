# htmlfind

Takes html on stdin, find things in it, and print it to stdout

## Install
```
go install github.com/emilen/htmlfind@latest
```

## Usage
```
Takes html on stdin, find things in it, and print it to stdout

Usage: htmlfind <command> [<args>]

Commands
        text <css-selector>     Return the text in the elements returned from the css-selector
        attr [<attribute>]      Return the value of the attributes
        comments                Return all HTML comments found in the documents
```
