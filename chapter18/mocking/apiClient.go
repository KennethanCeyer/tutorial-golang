package main


import (
    "errors"
    "net/http"
    "io/ioutil"
)


type APIClient interface {
    FetchData(endpoint string) (string, error)
}


type RealAPIClient struct{}


func (c *RealAPIClient) FetchData(endpoint string) (string, error) {
    resp, err := http.Get(endpoint)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()


    if resp.StatusCode != http.StatusOK {
        return "", errors.New("failed to fetch data")
    }


    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    return string(body), nil
}

