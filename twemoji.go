package emoji

import (
	"fmt"
	"strings"
)

const (
	TwemojiHTMLTemplate  = `<img src="%s" width="%d" height="%d" >`
	TwemojiXHTMLTemplate = `<img src="%s" width="%d" height="%d" />`
)

func UnicodeToTwemoji(src string, size int, isXHTML bool) string {
	for str, img := range str2img {
		var tpl string
		if isXHTML {
			tpl = TwemojiXHTMLTemplate
		} else {
			tpl = TwemojiHTMLTemplate
		}

		tag := fmt.Sprintf(tpl, img, size, size)

		src = strings.Replace(src, str, tag, -1)
	}

	return src
}

func EmojiTagToUnicode(src string) string {
	for name, chars := range name2codes {
		str := string(chars)
		tag := strings.Join([]string{`:`, name, `:`}, ``)

		src = strings.Replace(src, tag, str, -1)
	}

	return src
}
