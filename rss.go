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

var rssFeedUrls = []string{
	"http://gazeta.ru/export/rss/lenta.xml",
	"http://tvrain.ru/export/rss/programs/1018.xml",
	"http://interfax.ru/rss.asp",
	"https://buzzfeed.com/index.xml",
	"http://feeds.bbci.co.uk/news/world/rss.xml",
	"http://news.rambler.ru/rss/world/",
	"https://meduza.io/rss/all",
	"http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
}

var chats = map[int]struct{} {
	86082823: {},  // Igor
	162650098: {}, // Alisa
	89682072: {},  // Natalya
}

func userInterface()  {
	for true {
		resp := bot.GetUpdates(300)
		for e := range resp.Result {
			switch *resp.Result[e].Message.Text {
			case "/start":
				chats[resp.Result[e].Message.Chat.ID] = struct{}{}
			case "/stop":
				delete(chats, resp.Result[e].Message.Chat.ID)
			}
		}
	}
}

func main() {
	db, err := sql.Open("mysql", "GO_mysql_connector:L65gUIfd7i9JGHr4jhgH@(127.0.0.1:3306)/RSS_agregator_telegram_bot")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
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

	go userInterface()

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
					bot.SendNews(k, rssFeeds[e].Items[i].Title, rssFeeds[e].Items[i].Link)
				}
			}
			rssFeeds[e].Refresh = time.Now()
			rssFeeds[e].Items = nil
			rssFeeds[e].Unread = 0
		}
		time.Sleep(10*time.Second)
	}
}
