package util

import (
	"regexp"
)

func ParseTitle(resp string) string {
	receiveUrls := regexp.MustCompile(`<title>([\s\S]*?)</title>`).FindStringSubmatch(resp)
	if receiveUrls == nil {
		return "title not found."
	}

	return receiveUrls[1]
}

func ParseUrls(resp string) [][]string {
	receiveUrls := regexp.MustCompile("<a[\\s\\S]*?href=\"([\\s\\S]*?)\"[\\s\\S]*?>").FindAllStringSubmatch(resp, -1)
	return receiveUrls
}
