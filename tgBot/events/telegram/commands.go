package telegram

import (
	"log"
	"testbot/storage"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	log.Printf("got new command %s from %s", text, username)
	switch text {
	case RndCmd:
		return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)

	case StartCmd:
		return p.sendHello(chatID)

	default:
		return p.saveFilm(chatID, text, username)
	}
}

func (p *Processor) saveFilm(chatID int, name string, username string) error {
	film := &storage.Film{
		UserSended: username,
		FilmName:   name,
	}
	isExist, err := p.storage.IsExist(film)
	if err != nil {
		return err
	}
	if isExist {
		return p.tgClient.SendMessage(chatID, msgAlreadyExists)
	}
	if err := p.storage.Save(film); err != nil {
		return err
	}
	if err := p.tgClient.SendMessage(chatID, msgSaved); err != nil {
		return err
	}
	return nil
}

func (p *Processor) sendRandom(chatID int, username string) error {
	film, err := p.storage.PickRandom(username)
	if err != nil {
		return err
	}
	if p == nil {
		return p.tgClient.SendMessage(chatID, msgNoSavedFiles)
	}
	if err := p.tgClient.SendMessage(chatID, film.FilmName); err != nil {
		return err
	}
	return nil
}

func (p *Processor) sendHelp(chatID int) error {
	if err := p.tgClient.SendMessage(chatID, msgHelp); err != nil {
		return err
	}
	return nil
}

func (p *Processor) sendHello(chatID int) error {
	if err := p.tgClient.SendMessage(chatID, msgHello); err != nil {
		return err
	}
	return nil
}
