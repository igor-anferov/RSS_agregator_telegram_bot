package bd

import (
	"log"

	"github.com/jinzhu/gorm"
)

var bd *gorm.DB

type Feed struct {
	ID       uint   `gorm:"primary_key"`
	URL      string `gorm:"type:nvarchar(1024);unique_index;not null"`
	Standard bool
}

func (Feed) TableName() string {
	return "Feeds"
}

type User struct {
	gorm.Model
	ID uint `gorm:"primary_key"`
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

func (userFeed) TableName() string {
	return "User_Feeds"
}

func init() {
	db, err := gorm.Open("mysql", "GO_mysql_connector:L65gUIfd7i9JGHr4jhgH@(127.0.0.1:3306)/RSS_agregator_telegram_bot?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	bd = db
	bd.LogMode(true)
}

func Get() *gorm.DB {
	return bd
}

func CreateFeed(url string, fl bool) {
	bd.Create(&Feed{URL: url, Standard: fl})
}
