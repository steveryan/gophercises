package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	s := `<a href="#">
  Something here <a href="/dog">nested dog link</a>
</a>`

	tkn := html.NewTokenizer(strings.NewReader(s))
	links := make(map[string]string)
	currentLinks := make([]string, 0)
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
						currentLinks = append(currentLinks, s)
					}
				}
			}
		case html.TextToken:
			if len(currentLinks) > 0 {
				s := tkn.Token().Data
				s = strings.ReplaceAll(s, "\n", "")
				for _, v := range currentLinks {
					links[v] += s
				}
			}
		case html.EndTagToken:
			t := tkn.Token()
			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {
						s := a.Val
						s = strings.ReplaceAll(s, "\n", "")
						for i, v := range currentLinks {
							if v == s {
								currentLinks = append(currentLinks[:i], currentLinks[i+1:]...)
							}
						}
					}
				}
			}
		}
	}
	for k, v := range links {
		fmt.Printf("Href: %s\nText: %s\n\n", k, v)
	}
}
