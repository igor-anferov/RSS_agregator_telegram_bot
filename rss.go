package main

import (
	"fmt"
	"github.com/SlyMarbo/rss"
        "github.com/go-sql-driver/mysql"
)

func main() {

	result, err := rss.Fetch("https://meduza.io/rss/all")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
