package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := "https://www.calhoun.io"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("HTML:\n\n" + string(html) + "\n\n")

	links := getLinks(string(html))
	fmt.Println(strconv.Itoa(len(links)) + " links found")
}

func getLinks(s string) map[string]string {
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
	return links
}
