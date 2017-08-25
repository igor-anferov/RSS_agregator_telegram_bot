package bd

import (
	"log"

	"github.com/jinzhu/gorm"
)

var Bd *gorm.DB

type Feed struct {
	ID       uint   `gorm:"primary_key"`
	URL      string `gorm:"type:nvarchar(1024);unique_index;not null"`
	Standard bool
}

func (Feed) TableName() string {
	return "Feeds"
}

type User struct {
	ID int `gorm:"primary_key"`
}

func (User) TableName() string {
	return "Users"
}

type userFeed struct {
	UserID     uint   `gorm:"primary_key"`
	FeedID     uint   `gorm:"primary_key"`
	Desription string `gorm:"type:nvarchar(1024);unique_index;not null"`

	User User `gorm:"ForeignKey:UserID"`
	Feed Feed `gorm:"ForeignKey:FeedlID"`
}
type BotCommand struct {
	Command     string `gorm:"type:nchar(20);primary_key"`
	Description string `gorm:"type:nvarchar(1024);unique_index;not null"`
}

func (BotCommand) TableName() string {
	return "BotCommands"
}
func (userFeed) TableName() string {
	return "User_Feeds"
}

type SystemInfo struct {
	LastUpdateID int `gorm:"primary_key"`
}

func (SystemInfo) TableName() string {
	return "SystemInfo"
}
func init() {
	var err error
	Bd, err = gorm.Open("mysql", "GO_mysql_connector:L65gUIfd7i9JGHr4jhgH@(127.0.0.1:3306)/RSS_agregator_telegram_bot?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	Bd.LogMode(true)
}

func CreateFeed(url string, fl bool) {
	Bd.Create(&Feed{URL: url, Standard: fl})
}

func CreateUser(id int) {
	Bd.Create(&User{ID: id})
}

func GetUsersByFeedId(id int) []int {
	var idU []int
	Bd.Table("Users").Select("Users.id").Joins("JOIN User_Feeds JOIN Feeds ON Users.id = User_Feeds.`user` AND User_Feeds.feed = Feeds.id").Where("Feeds.id = ?", id).Pluck("Users.id", &idU)

	return idU
}

type UserFeeds struct {
	Url string
	Description *string
}

func GetFeedsByUserId(id int) []UserFeeds {
	var feeds []UserFeeds
	err := Bd.Table("Feeds").Joins("JOIN User_Feeds JOIN Users ON Users.id = User_Feeds.`user` AND User_Feeds.feed = Feeds.id").Where("Users.id = ?", id).Select("Feeds.url, User_Feeds.description").Find(&feeds).Error

	if err != nil {
		log.Println(err)
	}

	return feeds
}

func MyPluck() []string {
	var urlF []string
	err := Bd.Table("Feeds").Pluck("Feeds.url", &urlF).Error
	if err != nil {
		log.Fatal(err)
	}
	return urlF
}
