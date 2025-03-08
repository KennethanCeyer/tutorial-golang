package main

import (
	"io/ioutil"
	"net/http"
)


var ITNewsURIs = []string{
	"https://news.hada.io/",
	"https://news.hada.io/past",
	"https://news.hada.io/new",
}


type NewsLink struct {
	URI     string
	Title   string
	Content string
}


type NewsIndex struct {
	Link []NewsLink
}


func Indexing() (newsIndex NewsIndex, err error) {
	index := NewsIndex{[]NewsLink{}}
	for _, url := range ITNewsURIs {
		resp, err := http.Get(url)
		if err != nil {
			return NewsIndex{}, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		titleUrlPair := ParseIndex(body)
		for _, pair := range titleUrlPair {
			index.Link = append(index.Link, NewsLink{
				URI:   pair[0],
				Title: pair[1],
			})
		}
	}
	return index, nil
}
