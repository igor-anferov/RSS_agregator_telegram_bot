//bot_token := "441870254:AAHHaQbPt7abuqN97pD5nxGbtKhRIUUZGCI"
//api_url := "https://api.telegram.org/bot" + bot_token + "/"
package bot

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type client struct {
	Chat_id    int    `json:"chat_id"`    // свойство FirstName будет преобразовано в ключ "name"
	Text       string `json:"text"`       // свойство LastName будет преобразовано в ключ "lastname"
	Parse_mode string `json:"parse_mode"` // свойство Books будет преобразовано в ключ "ordered_books"
}

func SendMessageToIgor(mes string) {
	bot_token := "441870254:AAHHaQbPt7abuqN97pD5nxGbtKhRIUUZGCI"
	api_url := "https://api.telegram.org/bot" + bot_token + "/"
	Igor := &client{
		Chat_id:    86082823,
		Text:       mes,
		Parse_mode: "HTML",
	}

	//user1, _ := json.Marshal(Igor)
	//fmt.Println(string(user1))
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(Igor)
	http.Post(api_url+"sendMessage", "application/json", buf)
}
