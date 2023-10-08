package main

import (
  "fmt"
	"os"
	"encoding/json"
	"encoding/csv"
	"net/http"
	"strconv"
	"strings"
	"time"
	"os/signal"
	"syscall"
    twitterscraper "github.com/n0madic/twitter-scraper"
)

func main() {

	c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
        os.Exit(1)
    }()

    scraper := twitterscraper.New()
    
	f, _ := os.Open("cookies.json")
	// deserialize from JSON
	var cookies []*http.Cookie
	json.NewDecoder(f).Decode(&cookies)
	// load cookies
	scraper.SetCookies(cookies)
	// check login status
	scraper.IsLoggedIn()
	
	df, err := os.OpenFile("data1.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
       panic(err)
	}
	defer df.Close()

	// Create a CSV writer
	writer := csv.NewWriter(df)
	// defer writer.Flush()
	
	for {
		tweetids := []string{"1653786351228055552", "1543315251327901696", "1670168324322340864", "1663999865397952514", "1674360288445964288"}
		for _, s := range tweetids {
		
			tweet, err := scraper.GetTweet(s)
			if err != nil {
				panic(err)
			}
			fmt.Println(tweet.Likes)
			
			// Write a new row to the CSV file
			row := []string{time.Now().UTC().String(), s, strings.Replace(tweet.Text, "\n", "\\n", -1), strconv.Itoa(tweet.Likes)}
			err = writer.Write(row)
			if err != nil {
			   panic(err)
			}
			
		}
		writer.Flush()
		time.Sleep(10 * time.Minute) 
	}
}
