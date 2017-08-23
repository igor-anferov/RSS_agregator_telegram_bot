package main

import (
	"fmt"
	"log"
	"time"

	"github.com/SlyMarbo/rss"
	_ "github.com/go-sql-driver/mysql"
	"github.com/igor-anferov/RSS_agregator_telegram_bot/bd"
	"github.com/igor-anferov/RSS_agregator_telegram_bot/bot"
)

func main() {
	var database = bd.Get()
	defer database.Close()
	var f bd.Feed
	database.First(&f, 1)
	fmt.Println(f.URL)
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

	chats := []int{
		86082823,  // Igor
		162650098, // Alisa
		89682072,  // Natalya
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
				for k := range chats {
					bot.SendNews(chats[k], rssFeeds[e].Items[i].Title, rssFeeds[e].Items[i].Link)
				}
			}
			rssFeeds[e].Refresh = time.Now()
			rssFeeds[e].Items = nil
			rssFeeds[e].Unread = 0
		}
		time.Sleep(10 * time.Second)
	}
}
