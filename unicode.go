package emoji

import (
	"fmt"
	"strings"
)

func code2entities(src []rune) string {
	ret := make([]string, 0)

	for _, char := range src {
		ret = append(ret, fmt.Sprintf(`&#x%X;`, char))
	}

	return strings.Join(ret, ``)
}

func UnicodeToHTMLEntities(src string) string {
	for _, chars := range name2codes {
		str := string(chars)
		entities := code2entities(chars)

		src = strings.Replace(src, str, entities, -1)
	}

	return src
}
