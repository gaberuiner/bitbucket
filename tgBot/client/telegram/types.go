package telegram

type UpdatesResponse struct {
	OK      bool     `json:"ok"`
	Updates []Update `json:"result"`
}

type Update struct {
	ID      int             `json:"update_id"`
	Message IncomingMessage `json:"message"`
}

type IncomingMessage struct {
	From        From                `json:"User"`
	Chat        Chat                `json:"sender_chat"`
	Text        string              `json:"text"`
	ReplyMarkup ReplyKeyboardMarkup `json:"reply_markup"`
}

type ReplyKeyboardMarkup struct {
	Keyboard KeyboardButton `json:"keyboard"`
}

type KeyboardButton struct {
	Text map[string]string `json:"text"`
}

type From struct {
	Username string `json:"username"`
}

type Chat struct {
	Chat_id int `json:"id"`
}
