package telegram

import (
	"telegram-client"
	"e"
	"errors"
	"events"
	"storage"
	"strconv"
)

type IHandler interface {
	startBot(cmd, chatID string) error
	stopBot(chatID string) error

	sendHelp(chatID string) error
	sendGreeting(chatID string) error
	sendRandom(chatID string, userName string) error
	// SendRandomSet(chatID string, userName string, setLen int) error

	enableWarnings(cmd, chatID string) error
	disableWarnings(cmd, chatID string) error

	enableShortCuts(cmd, chatID string) error
	disableShortCuts(cmd, chatID string) error

	// setTimeInterval(cmd, chatID string) error

	// remove(cmd, chatID string) error
	// removeAll(cmd, chatID string) error

	// list(cmd, chatID string) error
	// timeInterval(cmd, chatID string) error

	// deleteHistory(chatID string)
}

type Processor struct {
	tg      *telegram.Client
	offset  int
	storage storage.IStorage
	handler IHandler
}

type Meta struct {
	ChatID   string `json:"chat_id"`
	UserName string `json:"username"`
}

var (
	ErrUnknownEventType = errors.New("unknown event type")
	ErrUnknownMetaType  = errors.New("unknown meta type")
)

func New(tg *telegram.Client, storage storage.IStorage) *Processor {
	return &Processor{
		tg:      tg,
		storage: storage,
	}
}

func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("failed to fetch updates", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}

	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil
}

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.processMessage(event)
	case events.Unknown:
		return e.Wrap("failed to process event", ErrUnknownEventType)
	}
	return nil
}

func (p *Processor) processMessage(event events.Event) (err error) {
	defer func() {
		err = e.Wrap("failed to process event", err)
	}()
	meta, err := meta(event)
	if err != nil {
		return err
	}
	if err := p.doCommand(event.Text, meta.ChatID, meta.UserName); err != nil {
		return err
	}
	return nil
}

func meta(event events.Event) (Meta, error) {
	res, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, e.Wrap("failed to cast meta", ErrUnknownMetaType)
	}
	return res, nil
}

func event(u telegram.Update) events.Event {
	updateType := fetchType(u)

	res := events.Event{
		Type: updateType,
		Text: fetchText(u),
	}

	if updateType == events.Message {
		res.Meta = Meta{
			ChatID:   strconv.Itoa(u.Message.Chat.ID),
			UserName: u.Message.From.Name,
		}
	}
	return res
}

func fetchText(u telegram.Update) string {
	if u.Message != nil {
		return ""
	}
	return u.Message.Text
}

func fetchType(u telegram.Update) events.Type {
	if u.Message == nil {
		return events.Unknown
	}
	return events.Message
}
