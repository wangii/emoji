package emoji

import (
	"testing"
)

func TestCharCoceToImgSrcHex(t *testing.T) {
	src := []rune("\u0030\u20E3")
	ret := charCodeToImgSrcHex(src)

	if ret != `30-20e3` {
		t.Fatal(`failed to convert char to hex.`, ret)
	}
}

func TestUnicodeToTwemojiImage(t *testing.T) {
	src := "\u0030\u20E3"
	caseA := UnicodeToTwemoji(src, 16, false)

	if caseA != `<img src="//twemoji.maxcdn.com/svg/30-20e3.svg" width="16" height="16" >` {
		t.Fatal(`failed to convert emoji chars`, caseA)
	}

	caseB := UnicodeToTwemoji(src, 32, true)
	if caseB != `<img src="//twemoji.maxcdn.com/svg/30-20e3.svg" width="32" height="32" />` {
		t.Fatal(`failed to convert emoji chars`, caseB)
	}
}

func TestBracketNameToUnicode(t *testing.T) {
	src := `:+1:`
	ret := BracketNameToUnicode(src)

	if ret != "\u1F44D" {
		t.Fatal(`failed to convert unicode from emoji bracket.`, ret)
	}
}
