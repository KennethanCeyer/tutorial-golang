package main

import (
	"regexp"
)


func ParseIndex(data []byte) [][]string {
	regexpString := `<div class=topictitle><a href='(https://[^']*)'[^>]*><h1>(.*?)</h1></a>`
	re := regexp.MustCompile(regexpString)
	content := string(data)
	matches := re.FindAllStringSubmatch(content, -1)
	urlTitlePair := make([][]string, len(matches))
	for i, value := range matches {
		urlTitlePair[i] = []string{value[1], value[2]}
	}
	return urlTitlePair
}
