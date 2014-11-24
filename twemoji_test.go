package emoji

import (
	"testing"
)

func TestUnicodeToTwemojiImage(t *testing.T) {
	src := string([]rune{0x30, 0xFE0F, 0x20E3})
	caseA := UnicodeToTwemoji(src, 16, false)

	if caseA != `<img src="//twemoji.maxcdn.com/svg/30-20e3.svg" width="16" height="16" >` {
		t.Fatal(`failed to convert emoji chars`, caseA)
	}

	caseB := UnicodeToTwemoji(src, 32, true)
	if caseB != `<img src="//twemoji.maxcdn.com/svg/30-20e3.svg" width="32" height="32" />` {
		t.Fatal(`failed to convert emoji chars`, caseB)
	}
}

func TestEmojiTagToUnicode(t *testing.T) {
	src := `:+1:`
	ret := EmojiTagToUnicode(src)

	if ret != string([]rune{0x1F44D}) {
		t.Fatal(`failed to convert unicode from emoji bracket.`, ret)
	}
}
