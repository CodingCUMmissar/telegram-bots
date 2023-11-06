package telegram

import (
	"e"
	"errors"
	"storage"
	"ternary"
)

func WarningsChangeState() {
	for _, cmd := range commandsWarnings {
		MutableCommands[cmd] = !MutableCommands[cmd]
	}
}

func (p *Processor) startBot(cmd, chatID string) error {
	if MutableCommands[cmd] {
		InitMutableCommands()
		MutableCommands[cmd] = false
		MutableCommands["IAmAbsolutelySureIWantToDeletemyHistory"] = true
		return p.sendGreeting(chatID)
	}
	return p.tg.SendMessage(chatID, msgBotAlreadyStarted)
}

func (p *Processor) stopBot(chatID string) error {
	DisableMostCommandsOnBotStop()
	return p.tg.SendMessage(chatID, msgBotStoppedWithTimeout)
}

func (p *Processor) enableWarnings(cmd, chatID string) error {
	WarningsChangeState()
	return p.tg.SendMessage(chatID, msgWarningsEnabled)
}

func (p *Processor) disableWarnings(cmd, chatID string) error {
	WarningsChangeState()
	return p.tg.SendMessage(chatID, msgWarningsEnabled)
}

func (p *Processor) enableShortCuts(cmd, chatID string) error {
	MutableCommands[cmd] = false
	MutableCommands["/disableshortcuts"] = true
	for _, cmd := range commandsShortcuts {
		MutableCommands[cmd] = true
	}
	return p.tg.SendMessage(chatID, msgShortcutsEnabled)
}

func (p *Processor) disableShortCuts(cmd, chatID string) error {
	MutableCommands[cmd] = false
	MutableCommands["/enableshortcuts"] = true
	for _, cmd := range commandsShortcuts {
		MutableCommands[cmd] = false
	}
	return p.tg.SendMessage(chatID, msgShortcutsDisabled)
}

// func (p *Processor) remove(cmd, chatID string) error {
// 	panic("not implemented")
// }

// func (p *Processor) removeAll(cmd, chatID string) error {
// 	panic("not implemented")
// }

// func (p *Processor) list(cmd, chatID string) error {
// 	panic("not implemented")
// }

// func (p *Processor) timeInterval(cmd, chatID string) error {
// 	panic("not implemented")
// }

func (p *Processor) sendHelp(chatID string) error {
	ok := MutableCommands["/warningsenabled"]
	msgHelp := ternary.Ternary[string](ok, msgHelpWithWarning, msgHelpNoWarning)

	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Processor) sendGreeting(chatID string) error {
	return p.tg.SendMessage(chatID, msgGreeting)
}

func (p *Processor) sendRandom(chatID string, userName string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: can't send random", err) }()
	sendMsg := NewMessageSender(chatID, p.tg)
	link, err := p.storage.PickRandom(userName)

	errIsNoSavedLinks := errors.Is(err, storage.ErrNoSavedLinks)
	if err != nil && !errIsNoSavedLinks {
		return
	}

	if errIsNoSavedLinks {
		return sendMsg(msgNoSavedLinks)
	}

	if err := sendMsg(link.URL); err != nil {
		return err
	}

	return p.storage.Remove(link)
}

// func (p *Processor) deleteHistory(chatID string) error {
// 	panic("not implemented")
// 	return p.tg.SendMessage(chatID, msgRemovedAll)
// }
