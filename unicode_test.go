package emoji

import (
	"testing"
)

func TestCharCoceToHTMLEntities(t *testing.T) {
	src := []rune("\u0030\u20E3")
	ret := charCodeToHTMLEntities(src)

	if ret != `&#x0030;&#x20E3;` {
		t.Fatal(`failed to convert chars to html entities`, ret)
	}
}

func TestUnicodeToHTMLEntities(t *testing.T) {
	src := "\u0030\u20E3"
	ret := UnicodeToHTMLEntities(src)

	if ret != `&#x0030;&#x20E3;` {
		t.Fatal(`failed to convert chars to entities`, ret)
	}
}
