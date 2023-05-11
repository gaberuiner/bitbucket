package events

import (
	"bot/client"
	storage "bot/database"
	"bot/funcs"
	"errors"
	"strings"
)

func (p *Processor) doCmd(username string, chatID int, message string) error {
	switch message {
	case "":
		return errors.New("message is empty")
	case "/start":
		p.tg.SendMessage(chatID, msgHello)
		if err := p.InlineButtonsStart(chatID, "Data Types, if you need help: /help"); err != nil {
			return err
		}

		return nil
	case "/help":
		p.tg.SendMessage(chatID, msgHelp)
		return nil
	default:
		data := splitString(message)
		if err := p.SaveSmtg(data, chatID, username); err != nil {
			return err
		}

	}

	return nil
}

func (p *Processor) SaveSmtg(data []string, chatID int, username string) error {
	if len(data) != 2 {
		if err := p.tg.SendMessage(chatID, msgUnknownCommand); err != nil {
			return err
		}
		return nil
	}
	switch data[0] {
	case "link":
		if funcs.IsURL(data[1]) {
			file := &storage.File{
				UserSended: username,
				Type:       data[0],
				FileName:   data[1],
			}
			if err := p.storage.Save(file); err != nil {
				return err
			} else {
				p.tg.SendMessage(chatID, msgSaved)
			}
		} else {
			p.tg.SendMessage(chatID, msgIsNotLink)
		}
	case "film":
		file := &storage.File{
			UserSended: username,
			Type:       data[0],
			FileName:   data[1],
		}
		if err := p.storage.Save(file); err != nil {
			return err
		} else {
			p.tg.SendMessage(chatID, msgSaved)
		}
	case "book":
		file := &storage.File{
			UserSended: username,
			Type:       data[0],
			FileName:   data[1],
		}
		if err := p.storage.Save(file); err != nil {
			return err
		} else {
			p.tg.SendMessage(chatID, msgSaved)
		}

	case "image":
		if funcs.IsImageURL(data[1]) {
			file := &storage.File{
				UserSended: username,
				Type:       data[0],
				FileName:   data[1],
			}
			if err := p.storage.Save(file); err != nil {
				return err
			} else {
				p.tg.SendMessage(chatID, msgSaved)
			}
		} else {
			p.tg.SendMessage(chatID, msgIsNotImageURL)
		}
	}
	return nil
}

func (p *Processor) saveFile(chatID int, name string, username string, Qtype string) error {
	film := &storage.File{
		UserSended: username,
		Type:       Qtype,
		FileName:   name,
	}
	isExist, err := p.storage.IsExist(film)
	if err != nil {
		return err
	}
	if isExist {
		return p.tg.SendMessage(chatID, msgAlreadyExists)
	}
	if err := p.storage.Save(film); err != nil {
		return err
	}
	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {
		return err
	}
	return nil
}

func (p *Processor) sendRandom(chatID int, username string, Qtype string) error {
	file, err := p.storage.PickRandom(Qtype)
	if err != nil {
		return err
	}
	if file == "" {
		return p.tg.SendMessage(chatID, msgNoSavedFiles)
	}
	if err := p.tg.SendMessage(chatID, file); err != nil {
		return err
	}
	return nil
}

func (p *Processor) sendALL(chatID int, username string, Qtype string) error {
	if Qtype != "link" && Qtype != "image" {
		file, err := p.storage.SelectALL(Qtype)
		if err != nil {
			return err
		}
		if file == "" {
			return p.tg.SendMessage(chatID, msgNoSavedFiles)
		}
		if err := p.tg.SendMessage(chatID, file); err != nil {
			return err
		}
		return nil
	} else {
		file, err := p.storage.SelectALL(Qtype)
		if err != nil {
			return err
		}
		if file == "" {
			return p.tg.SendMessage(chatID, msgNoSavedFiles)
		}
		files := strings.Fields(file)
		for _, q := range files {
			if err := p.tg.SendMessage(chatID, q); err != nil {
				return err
			}
		}

		return nil
	}
}

func (p *Processor) sendHelp(chatID int) error {
	if err := p.tg.SendMessage(chatID, msgHelp); err != nil {
		return err
	}
	return nil
}

func (p *Processor) sendHello(chatID int) error {
	if err := p.tg.SendMessage(chatID, msgHello); err != nil {
		return err
	}
	return nil
}

func (p *Processor) InlineButtonsStart(chatID int, message string) error {
	inlineButton := [][]client.InlineKeyboardButton{
		{
			{
				Text:         "Learn",
				CallbackData: "button1",
			},
			{
				Text:         "Film",
				CallbackData: "button2",
			},
			{
				Text:         "Book",
				CallbackData: "button3",
			},
			{
				Text:         "Image",
				CallbackData: "button4",
			},
		},
	}

	if err := p.tg.SendInlineButtons(chatID, inlineButton, message); err != nil {
		return err
	}
	return nil
}

func splitString(s string) []string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return []string{}
	}
	words[0] = strings.ToLower(words[0])

	return []string{words[0], strings.Join(words[1:], " ")}
}
