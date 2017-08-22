package main

import (
	"fmt"
	"log"
	"time"

	"github.com/SlyMarbo/rss"
	_ "github.com/go-sql-driver/mysql"
	"github.com/igor-anferov/RSS_agregator_telegram_bot/bot"
	"github.com/jinzhu/gorm"
)

type feed struct {
	gorm.Model

	ID       uint   `gorm:"primary_key"`
	URL      string `gorm:"type:nvarchar(1024);unique_index;not null"`
	Standard bool
}

func (feed) TableName() string {
	return "feeds"
}

type user struct {
	ID uint `gorm:"primary_key"`
}

func (user) TableName() string {
	return "users"
}

type userFeed struct {
	UserID     uint   `gorm:"not null"`
	FeedID     uint   `gorm:"not null"`
	Desription string `gorm:"type:nvarchar(1024);unique_index;not null"`

	User user `gorm:"ForeignKey:UserID"`
	Feed feed `gorm:"ForeignKey:FeedlID"`
	ID   uint `gorm:"primary_key(UserID, FeedID)"`
}

func (userFeed) TableName() string {
	return "users_feeds"
}

func main() {
	db, err := gorm.Open("mysql", "GO_mysql_connector:L65gUIfd7i9JGHr4jhgH@(127.0.0.1:3306)/RSS_agregator_telegram_bot?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	/*if err := db.Ping(); err != nil {
		log.Fatal(err)
	}*/
	defer db.Close()

	db.LogMode(true)

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
