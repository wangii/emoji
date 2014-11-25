package emoji

import (
	"testing"
)

func TestUnicodeToHTMLEntities(t *testing.T) {
	src := string([]rune{0x30, 0xFE0F, 0x20E3})

	if ret := UnicodeToHTMLEntities(src); ret != `&#x30;&#xFE0F;&#x20E3;` {
		t.Fatal(`failed to convert emoji unicode to html entities.`, ret)
	}
}
