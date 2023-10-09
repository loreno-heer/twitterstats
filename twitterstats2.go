package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
	"flag"

	twitterscraper "github.com/n0madic/twitter-scraper"
)




func main() {

	var tweet_ids []string
	var out_file string
	var cookie_file string

	flag.StringVar(&out_file, "o", "", "output file")
	flag.StringVar(&cookie_file, "c", "cookies.json", "cookie file")
	tweet_ids = flag.Args()
	
	flag.Parse()
	if len(out_file) == 0 || len(tweet_ids) == 0 {
		fmt.Println("Usage: twitterstats -o output.csv -c cookies.json tweetid1 [tweetid2 ...]")
		flag.PrintDefaults()
        os.Exit(1)
	}
	

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Exit(0)
	}()

	scraper := twitterscraper.New()

	f, _ := os.Open(cookie_file)
	// deserialize from JSON
	var cookies []*http.Cookie
	json.NewDecoder(f).Decode(&cookies)
	// load cookies
	scraper.SetCookies(cookies)
	// check login status
	scraper.IsLoggedIn()

	df, err := os.OpenFile(out_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer df.Close()

	// Create a CSV writer
	writer := csv.NewWriter(df)
	// defer writer.Flush()

	for {
		for _, s := range tweet_ids {

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
		time.Sleep(5 * time.Minute)
	}
}
