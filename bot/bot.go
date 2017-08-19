//bot_token := "441870254:AAHHaQbPt7abuqN97pD5nxGbtKhRIUUZGCI"
//api_url := "https://api.telegram.org/bot" + bot_token + "/"
package bot

import (
	"bytes"
	"encoding/json"
	"net/http"
	"fmt"
)

type Button struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

type ButtonsGrid struct {
	Inline_keyboard [][]Button `json:"inline_keyboard"`
}

type message struct {
	Chat_id    int    `json:"chat_id"`    // свойство FirstName будет преобразовано в ключ "name"
	Text       string `json:"text"`       // свойство LastName будет преобразовано в ключ "lastname"
	Parse_mode string `json:"parse_mode"` // свойство Books будет преобразовано в ключ "ordered_books"
	Reply_markup ButtonsGrid `json:"reply_markup"`
}

func SendNews(chat_id int, title string, url string) {
	bot_token := "441870254:AAHHaQbPt7abuqN97pD5nxGbtKhRIUUZGCI"
	api_url := "https://api.telegram.org/bot" + bot_token + "/"

	newsMessage := message{
		chat_id,
		"<a href=\"" + url + "\">" + title + "</a>",
		"HTML",
		ButtonsGrid{
			[][]Button {
				{
					{
						"Смотреть подробнее",
						url,
					},
				},
			},
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(newsMessage)
	fmt.Println(buf)
	http.Post(api_url+"sendMessage", "application/json", buf)
}
