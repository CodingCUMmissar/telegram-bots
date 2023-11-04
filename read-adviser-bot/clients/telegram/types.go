package telegram

type Update struct {
	ID      int    `json:"update_id"`
	Message string `json:"message"`
}

type UpdatesReponse struct {
	Ok      bool     `json:"ok"`
	Results []Update `json:"results"`
}
