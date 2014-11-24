package emoji

import (
	"fmt"
	"strings"
)

func charCodeToImgSrcHex(src []rune) string {
	buf := make([]string, 0)

	for _, char := range src {
		buf = append(buf, fmt.Sprintf(`%x`, char))
	}

	return strings.Join(buf, `-`)
}

func UnicodeToTwemoji(src string, size int, isXHTML bool) string {
	for char, imgSrc := range char2imgSrc {
		tag := fmt.Sprintf(
			`<img src="%s" width="%d" height="%d" >`,
			imgSrc,
			size, size)

		if isXHTML {
			tag = strings.Replace(tag, ` >`, ` />`, 1)
		}

		src = strings.Replace(src, char, tag, -1)
	}

	return src
}

func BracketNameToUnicode(src string) string {
	for name, char := range name2char {
		src = strings.Replace(src, fmt.Sprintf(`:%s:`, name), char, -1)
	}

	return src
}
