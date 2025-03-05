package main

import (
	"io/ioutil"
	"net/http"
	"sync"
)


const UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6)" +
	" AppleWebKit/537.36 (KHTML, like Gecko)" +
	" Chrome/79.0.3945.130 Safari/537.36"


func Crawler(index NewsIndex, useGoroutine bool) error {
	var wg sync.WaitGroup
	for i, link := range index.Link {
		wg.Add(1)
		if useGoroutine {
			go GetContent(link.URI, &index.Link[i], &wg)
		} else {
			err := GetContent(link.URI, &index.Link[i], &wg)
			if err != nil {
				return err
			}
		}
	}
	wg.Wait()
	return nil
}


func GetContent(
	url string, link *NewsLink, wg *sync.WaitGroup) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Close = true
	req.Header.Add("User-Agent", UserAgent)
	resp, err := client.Do(req)
	if err != nil {
		wg.Done()
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		wg.Done()
		return err
	}
	link.Content = string(body)
	wg.Done()
	return nil
}
