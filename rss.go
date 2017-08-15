package main

import (
	"fmt"

	"github.com/SlyMarbo/rss"
)

func main() {

	result, err := rss.Fetch("https://meduza.io/rss/all")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
