package telegram

type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

type UpdatesReponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"results"`
}

type IncomingMessage struct {
	Text string `json:"text"`
	From *User  `json:"from,omitempty"`
	Chat *Chat  `json:"chat,omitempty"`
}

type Chat struct {
	ID int `json:"id"`
}

type User struct {
	Name string `json:"username"`
}
