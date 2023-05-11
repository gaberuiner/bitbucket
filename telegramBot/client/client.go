package client

import (
	"bot/errors"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type Client struct {
	host     string
	basePath string

	client http.Client
}

var (
	errorRequest      = "can't make request"
	errorUpdates      = "can't get updates"
	errorUnmarshal    = "can't unmarshal data"
	errorSendMessage  = "can't send message"
	getUpdatesMethod  = "getUpdates"
	sendMessageMethod = "sendMessage"
)

func New(host string, token string) Client {
	return Client{
		host:     host,
		basePath: basePath(token),
		client:   http.Client{},
	}
}

func basePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))
	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, errors.Error(errorUpdates, err)
	}
	var res GetUpdate
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, errors.Error(errorUnmarshal, err)
	}

	for _, update := range res.Result {
		if update.Message.NewChatMembers != nil {
			for _, member := range update.Message.NewChatMembers {
				err := c.SendMessage(update.Message.From.UserID, "Hello "+member.Username+"!")
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return res.Result, nil
}

func (c *Client) SendMessage(chatID int, text string) error {
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", text)
	_, err := c.doRequest(sendMessageMethod, q)
	if err != nil {
		return errors.Error(errorSendMessage, err)
	}
	return nil
}

func (c *Client) doRequest(method string, query url.Values) ([]byte, error) {
	u := url.URL{
		Scheme: "https",
		Host:   c.host,
		Path:   path.Join(c.basePath, method),
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Error(errorRequest, err)
	}
	req.URL.RawQuery = query.Encode()
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Error(errorRequest, err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Error(errorRequest, err)
	}
	return body, nil
}

func (c *Client) SendInlineButtons(chatID int, inlineButton [][]InlineKeyboardButton, message string) error {
	inlineKeyboardMarkup := InlineKeyboardMarkup{
		InlineKeyboard: inlineButton,
	}
	q := url.Values{}
	q.Add("chat_id", strconv.Itoa(chatID))
	q.Add("text", message)
	inlineKeyboard, err := json.Marshal(inlineKeyboardMarkup)
	if err != nil {
		return errors.Error(errorSendMessage, err)
	}
	q.Add("reply_markup", string(inlineKeyboard))
	_, err = c.doRequest(sendMessageMethod, q)
	if err != nil {
		return errors.Error(errorSendMessage, err)
	}

	return nil
}
