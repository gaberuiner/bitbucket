package events

import (
	"bot/client"
)

var (
	noSavedData     = "no saved data"
	noDataInStorage = "no data in storage"
)

func (p *Processor) InlineButtonProcess(username string, chatID int, button string, message string) error {
	switch button {
	case "button1":

		inlineButton := [][]client.InlineKeyboardButton{
			{
				{
					Text:         "Random Link",
					CallbackData: "link1",
				},
				{
					Text:         "All Links",
					CallbackData: "link2",
				},
			},
		}
		if err := p.tg.SendInlineButtons(chatID, inlineButton, "choose"); err != nil {
			return err
		}
		return nil
	case "button2":
		inlineButton := [][]client.InlineKeyboardButton{
			{
				{
					Text:         "Random Film",
					CallbackData: "film1",
				},
				{
					Text:         "All Films",
					CallbackData: "film2",
				},
			},
		}
		if err := p.tg.SendInlineButtons(chatID, inlineButton, "choose"); err != nil {
			return err
		}
		return nil

	case "button3":
		inlineButton := [][]client.InlineKeyboardButton{
			{
				{
					Text:         "Random Book",
					CallbackData: "book1",
				},
				{
					Text:         "All Books",
					CallbackData: "book2",
				},
			},
		}
		if err := p.tg.SendInlineButtons(chatID, inlineButton, "choose"); err != nil {
			return err
		}
		return nil

	case "button4":
		inlineButton := [][]client.InlineKeyboardButton{
			{
				{
					Text:         "Random Img",
					CallbackData: "img1",
				},
				{
					Text:         "All Imgs",
					CallbackData: "img2",
				},
			},
		}
		if err := p.tg.SendInlineButtons(chatID, inlineButton, "choose"); err != nil {
			return err
		}
		return nil
	case "link1":
		err := p.sendRandom(chatID, username, "link")
		if err != nil {
			if err := p.tg.SendMessage(chatID, noDataInStorage); err != nil {
				return err
			}
		}
		return nil
	case "link2":
		err := p.sendALL(chatID, username, "link")
		if err != nil {
			if err := p.tg.SendMessage(chatID, noDataInStorage); err != nil {
				return err
			}
		}
		return nil
	case "film1":
		err := p.sendRandom(chatID, username, "film")
		if err != nil {
			if err := p.tg.SendMessage(chatID, noDataInStorage); err != nil {
				return err
			}
		}
		return nil
	case "film2":
		err := p.sendALL(chatID, username, "film")
		if err != nil {
			if err := p.tg.SendMessage(chatID, noDataInStorage); err != nil {
				return err
			}
		}
		return nil
	case "book1":
		err := p.sendRandom(chatID, username, "book")
		if err != nil {
			if err := p.tg.SendMessage(chatID, noDataInStorage); err != nil {
				return err
			}
		}
		return nil
	case "book2":
		err := p.sendALL(chatID, username, "book")
		if err != nil {
			if err := p.tg.SendMessage(chatID, noDataInStorage); err != nil {
				return err
			}
		}
		return nil
	case "img1":

		err := p.sendRandom(chatID, username, "image")
		if err != nil {
			if err := p.tg.SendMessage(chatID, noDataInStorage); err != nil {
				return err
			}
		}
		return nil
	case "img2":

		err := p.sendALL(chatID, username, "image")
		if err != nil {
			if err := p.tg.SendMessage(chatID, noDataInStorage); err != nil {
				return err
			}
		}
		return nil

	}

	return nil
}
