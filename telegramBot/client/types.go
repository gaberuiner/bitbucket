package client

type GetUpdate struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID            int           `json:"update_id"`
	Message       Message       `json:"message"`
	CallbackQuery callbackQuery `json:"callback_query"`
}

type callbackQuery struct {
	ID              string  `json:"id"`
	From            From    `json:"from"`
	Message         Message `json:"message"`
	InlineMessageID string  `json:"inline_message_id"`
	Data            string  `json:"data"`
}

type Message struct {
	MessageID      int    `json:"message_id"`
	From           From   `json:"from"`
	Text           string `json:"text"`
	NewChatMembers []From `json:"new_chat_members"`
}

type From struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}
