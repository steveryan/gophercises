package main

import (
	"reflect"
	"testing"
)

func TestGetLinksWithNoHTML(t *testing.T) {
	result := getLinks("")

	if len(result) != 0 {
		t.Errorf("Result was incorrect, got: %s, want: nil.", result)
	}
}

func TestGetLinksWithNoLinks(t *testing.T) {
	result := getLinks("Some text")

	if len(result) != 0 {
		t.Errorf("Result was incorrect, got: %s, want: nil.", result)
	}
}

func TestGetLinksWithOneLink(t *testing.T) {
	result := getLinks(`<a href="#">Some text</a>`)
	expected := map[string]string{"#": "Some text"}
	equal := reflect.DeepEqual(result, expected)
	if !equal {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result, expected)
	}
}
