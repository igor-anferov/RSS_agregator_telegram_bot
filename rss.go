package main

import (
	"github.com/SlyMarbo/rss"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	"time"
	"github.com/igor-anferov/RSS_agregator_telegram_bot/bot"
	"log"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "GO_mysql_connector:L65gUIfd7i9JGHr4jhgH@(127.0.0.1:3306)/RSS_agregator_telegram_bot")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	rssFeedUrls := []string{
		"http://gazeta.ru/export/rss/lenta.xml",
		"http://tvrain.ru/export/rss/programs/1018.xml",
		"http://interfax.ru/rss.asp",
		"https://buzzfeed.com/index.xml",
		"http://feeds.bbci.co.uk/news/world/rss.xml",
		"http://news.rambler.ru/rss/world/",
		"https://meduza.io/rss/all",
		"http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
	}

	chats := []int {
		86082823, // Igor
		162650098, // Alisa
	}

	rssFeeds := make([]*rss.Feed, 0, len(rssFeedUrls))
	for e := range rssFeedUrls {
		result, err := rss.Fetch(rssFeedUrls[e])
		if err != nil {
			log.Fatal(err)
		}
		result.Refresh = time.Now()
		result.Items = nil
		result.Unread = 0
		rssFeeds = append(rssFeeds, result)
	}

	for true {
		for e := range rssFeeds {
			err := rssFeeds[e].Update()
			if err != nil {
				log.Print(err)
				log.Println(rssFeeds[e].Title)
			}
			for i := range rssFeeds[e].Items {
				fmt.Println(rssFeeds[e].Items[i])
				for e := range chats {
					bot.SendNews(chats[e], rssFeeds[e].Items[i].Title, rssFeeds[e].Items[i].Link)
				}
			}
			rssFeeds[e].Refresh = time.Now()
			rssFeeds[e].Items = nil
			rssFeeds[e].Unread = 0
		}
		time.Sleep(10*time.Second)
	}
}
