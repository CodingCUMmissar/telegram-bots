package telegram

import (
	"e"
	"errors"
	"log"
	"slices"
	"storage"
	"strings"
)

func (p *Processor) doCommand(text, chatID, userName string) error {
	text = strings.TrimSpace(text)
	sendMsg := NewMessageSender(chatID, p.tg)
	log.Printf("new command '%s' from '%s'", text, userName)

	if isAddCommand(text) {
		return p.savePage(text, chatID, userName)
	}

	slices.Sort[[]string](commandsFull)

	if ok := slices.Contains[[]string, string](commandsFull, text); !ok {
		return sendMsg(msgUnknownCommand)
	}

	return p.handleCommand(text, userName, chatID)

}

func (p *Processor) handleCommand(cmd, userName, chatID string) error {

	switch cmd {
	case "/start":
		return p.startBot(cmd, chatID)

	case "/stop":
		return p.stopBot(chatID)

	case "/help":
	case "/h":
		return p.sendHelp(cmd)

	case "/greeting":
	case "/g":
		return p.handler.sendGreeting(chatID)

	case "/random":
	case "/rnd":
		return p.handler.sendRandom(chatID, userName)

	case "/disablewarning":
		return p.handler.disableWarnings(cmd, chatID)

	case "/enablewarning":
		return p.handler.enableWarnings(cmd, chatID)

	case "/enableshortcuts":
		return p.handler.enableShortCuts(cmd, chatID)

	case "/disableshortcuts":
		return p.handler.disableShortCuts(cmd, chatID)

	// case "/settime":
	// case "/sett":
	// 	return p.handler.SetTimeInterval(cmd, chatID)

	// case "/remove %s":
	// case "/rm %s":
	// 	return p.handler.Remove(cmd, chatID)

	// case "/removeall":
	// case "/rmall":
	// 	return p.handler.RemoveAll(cmd, chatID)

	// case "/list":
	// case "/l":
	// 	return p.handler.List(cmd, chatID)

	// case "/time":
	// case "/t":
	// 	return p.handler.TimeInterval(cmd, chatID)

	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}

	return errors.New("impossible???????")
}

func (p *Processor) savePage(pageURL, chatID, userName string) (err error) {
	defer func() {
		err = e.WrapIfErr("failed to save page", err)
	}()

	sendMsg := NewMessageSender(chatID, p.tg)

	page := &storage.Page{
		URL:      pageURL,
		UserName: userName,
	}

	isExists, err := p.storage.IsExists(page)
	if err != nil {
		return err
	}

	if isExists {
		return sendMsg(msgLinkExists)
	}

	if err := p.storage.Save(page); err != nil {
		return err
	}

	if err := sendMsg(msgLinkSaved); err != nil {
		return err
	}

	return nil
}
