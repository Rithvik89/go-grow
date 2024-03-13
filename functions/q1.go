package functions

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func visit(n *html.Node, stack []string) []string {

	if n.Type == html.ElementNode && n.Data == "a" {
		// iterate over the node attr
		for _, a := range n.Attr {
			if a.Key == "href" {
				stack = append(stack, a.Val)
			}
		}
	}

	for i := n.FirstChild; i != nil; i = i.NextSibling {
		visit(i, stack)
	}

	return stack
}

func main() {
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Findlinks error: %v\n", err)
		os.Exit(1)
	}

	// get string of links into link string[]
	var stack []string

	links := visit(doc, stack)
	for link := range links {
		fmt.Printf("Link is %v\n", link)
	}
}
