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

var chats = map[int]struct{}{
	86082823:  {}, // Igor
	162650098: {}, // Alisa
	89682072:  {}, // Natalya
}

func userInterface() {
	for true {
		resp := bot.GetUpdates(300)
		for e := range resp.Result {
			switch *resp.Result[e].Message.Text {
			case "/start":
				bd.CreateUser(resp.Result[e].Message.Chat.ID)
			case "/myfeeds":
				feeds := bd.GetFeedsByUserId(resp.Result[e].Message.Chat.ID)
				text := "<b>🗞 Your feeds:</b>\n"
				for e := range feeds {
					if feeds[e].Description == nil {
						result, err := rss.Fetch(feeds[e].Url)
						if err != nil {
							log.Println(err)
							continue
						}
						feeds[e].Description = &result.Title
					}
					text += "▪️ " + *feeds[e].Description + "\n"
				}
				bot.SendMessage(resp.Result[e].Message.Chat.ID, text)
			default:
				bot.SendMessage(resp.Result[e].Message.Chat.ID, "Sorry, I don't understood you")
			}
		}
	}
}

func main() {
	defer bd.Bd.Close()

	var rssFeedUrls []string
	rssFeedUrls = bd.MyPluck()

	fmt.Println(rssFeedUrls[1])
	//fmt.Println(bd.Select(7))

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
		time.Sleep(10 * time.Second)
	}
}
