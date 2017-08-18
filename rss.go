package main

import (
	"github.com/SlyMarbo/rss"
	//"github.com/go-sql-driver/mysql"

	"time"

	"github.com/igor-anferov/RSS_agregator_telegram_bot/bot"
)

func main() {
	rssFeedUrls := []string{
		"http://gazeta.ru/export/rss/lenta.xml",
		"https://tvrain.ru/export/rss/all.xml",
		"http://interfax.ru/rss.asp",
		"https://buzzfeed.com/index.xml",
		"http://feeds.bbci.co.uk/news/world/rss.xml",
		"http://news.rambler.ru/rss/world/",
		"https://meduza.io/rss/all",
		"http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml",
	}

	rssFeeds := make([]*rss.Feed, 0, len(rssFeedUrls))
	for e := range rssFeedUrls {
		result, err := rss.Fetch(rssFeedUrls[e])
		if err != nil {
			panic(err)
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
				panic(err)
			}
			for i := range rssFeeds[e].Items {
				//fmt.Println(rssFeeds[e].Items[i])
				//res, _ := json.Marshal(rssFeeds[e].Items[i])
				bot.SendMessageToIgor(rssFeeds[e].Items[i].Link)
			}
			rssFeeds[e].Refresh = time.Now()
			rssFeeds[e].Items = nil
			rssFeeds[e].Unread = 0
		}
	}
}
