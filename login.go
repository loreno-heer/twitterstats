package main

import (
    "context"
    "fmt"
    "os"
	  "encoding/json"
    twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {
    scraper := twitterscraper.New()
    err := scraper.Login("username", "password")
    if err != nil {
        panic(err)
    }
    cookies := scraper.GetCookies()
    // serialize to JSON
    js, _ := json.Marshal(cookies)
    // save to file
    f, _ = os.Create("cookies.json")
    f.Write(js)
}
