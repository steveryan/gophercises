package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	s := `<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>`

	tkn := html.NewTokenizer(strings.NewReader(s))
	links := make(map[string]string)
	var currentLink string
loop:
	for {
		token := tkn.Next()
		switch token {
		case html.ErrorToken:
			break loop
		case html.StartTagToken:
			t := tkn.Token()
			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {
						s := a.Val
						s = strings.ReplaceAll(s, "\n", "")
						links[s] = ""
						currentLink = s
					}
				}
			}
		case html.TextToken:
			if currentLink != "" {
				s := tkn.Token().Data
				s = strings.ReplaceAll(s, "\n", "")
				links[currentLink] += s
			}
		case html.EndTagToken:
			if tkn.Token().Data == "a" {
				currentLink = ""
			}
		}
	}
	for k, v := range links {
		fmt.Printf("Href: %s\nText: %s\n\n", k, v)
	}
}
