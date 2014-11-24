package emoji

import (
	"fmt"
	"strings"
)

func charCodeToHTMLEntities(src []rune) string {
	buf := make([]string, 0)
	for _, char := range src {
		buf = append(buf, fmt.Sprintf(`&#x%04X;`, char))
	}

	return strings.Join(buf, ``)
}

func UnicodeToHTMLEntities(src string) string {
	for _, char := range name2char {
		src = strings.Replace(src, char, charCodeToHTMLEntities([]rune(char)), -1)
	}

	return src
}
