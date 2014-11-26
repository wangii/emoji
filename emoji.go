package emoji

import (
	"fmt"
	"strings"
)

const (
	TwemojiHTMLTemplate  = `<img src="%s" width="%d" height="%d" >`
	TwemojiXHTMLTemplate = `<img src="%s" width="%d" height="%d" />`
)

func code2entities(src []rune) string {
	ret := make([]string, 0)

	for _, char := range src {
		ret = append(ret, fmt.Sprintf(`&#x%X;`, char))
	}

	return strings.Join(ret, ``)
}

func EmojiTagToHTMLEntities(src string) string {
	for name, chars := range name2codes {
		tag := strings.Join([]string{`:`, name, `:`}, ``)

		src = strings.Replace(src, tag, code2entities(chars), -1)
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

func EmojiTagToTwemoji(src string, size int, isXHTML bool) string {
	for name, chars := range name2codes {
		str := string(chars)
		if img, ok := str2img[str]; ok {
			var tpl string
			if isXHTML {
				tpl = TwemojiXHTMLTemplate
			} else {
				tpl = TwemojiHTMLTemplate
			}

			imgTag := fmt.Sprintf(tpl, img, size, size)
			tagStr := strings.Join([]string{`:`, name, `:`}, ``)

			src = strings.Replace(src, tagStr, imgTag, -1)
		}
	}

	return src
}

func UnicodeToHTMLEntities(src string) string {
	for _, chars := range name2codes {
		str := string(chars)
		entities := code2entities(chars)

		src = strings.Replace(src, str, entities, -1)

		chars2 := make([]rune, 0)
		for _, char := range chars {
			if char == '\uFE0F' {
				continue
			}

			chars2 = append(chars2, char)
		}

		str = string(chars2)
		entities = code2entities(chars2)

		src = strings.Replace(src, str, entities, -1)
	}

	return src
}

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
