package bot

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/igor-anferov/RSS_agregator_telegram_bot/bd"
)

var bot_token = "441870254:AAHHaQbPt7abuqN97pD5nxGbtKhRIUUZGCI"
var api_url = "https://api.telegram.org/bot" + bot_token + "/"

type Button struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

type ButtonsGrid struct {
	Inline_keyboard [][]Button `json:"inline_keyboard"`
}

type SendMessageReq struct {
	Chat_id      int         `json:"chat_id"`    // свойство FirstName будет преобразовано в ключ "name"
	Text         string      `json:"text"`       // свойство LastName будет преобразовано в ключ "lastname"
	Parse_mode   string      `json:"parse_mode"` // свойство Books будет преобразовано в ключ "ordered_books"
	Reply_markup ButtonsGrid `json:"reply_markup"`
}

type User struct {
	ID            int     `json:"id"`
	Is_bot        bool    `json:"is_bot"`
	First_name    string  `json:"first_name"`
	Last_name     *string `json:"last_name"`
	Username      *string `json:"username"`
	Language_code *string `json:"language_code"`
}

type Chat struct {
	ID int `json:"id"`
}

type Message struct {
	Message_id int     `json:"message_id"`
	From       *User   `json:"from"`
	Chat       Chat    `json:"chat"`
	Text       *string `json:"text"`
}

type InlineQuery struct {
	ID     string `json:"id"`
	From   User   `json:"from"`
	Query  string `json:"query"`
	Offset string `json:"offset"`
}

type Update struct {
	Update_id           int          `json:"update_id"`
	Message             *Message     `json:"message"`
	Edited_message      *Message     `json:"edited_message"`
	Channel_post        *Message     `json:"channel_post"`
	Edited_channel_post *Message     `json:"edited_channel_post"`
	Inline_query        *InlineQuery `json:"inline_query"`
}

type GetUpdatesResponse struct {
	Ok          bool     `json:"ok"`
	Error_code  int      `json:"error_code"`
	Description string   `json:"description"`
	Result      []Update `json:"result"`
}

func GetUpdates(timeout int) GetUpdatesResponse {
	var update []int
	err := bd.Bd.Table("SystemInfo").Pluck("SystemInfo.lastUpdateId", &update).Error
	if err != nil {
		log.Fatal(err)
	}
	getUpdtsReq := struct {
		Offset  int `json:"offset"`
		Limit   int `json:"limit"`
		Timeout int `json:"timeout"`
	}{
		update[0] + 1,
		1,
		timeout,
	}
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(getUpdtsReq)
	resp, err := http.Post(api_url+"getUpdates", "application/json", buf)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	var response GetUpdatesResponse
	err = json.Unmarshal([]byte(respBody), &response)
	if err != nil {
		log.Fatal(err)
	}
	if !response.Ok {
		log.Print(response.Error_code, response.Description)
	}
	if len(response.Result) > 0 {
		bd.Bd.Table("SystemInfo").Update("lastUpdateId", response.Result[0].Update_id)
	}
	return response
}

func SendNews(chat_id int, title string, url string) {
	newsMessage := SendMessageReq{
		chat_id,
		"<a href=\"" + url + "\">" + title + "</a>",
		"HTML",
		ButtonsGrid{
			[][]Button{
				{
					{
						"☝🏻  Смотреть подробнее  👀",
						url,
					},
				},
			},
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(newsMessage)
	http.Post(api_url+"sendMessage", "application/json", buf)
}

func SendMessage(chat_id int, text string) {

	newsMessage := SendMessageReq{
		chat_id,
		text,
		"HTML",
		ButtonsGrid{
			[][]Button{},
		},
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(newsMessage)
	http.Post(api_url+"sendMessage", "application/json", buf)
}
