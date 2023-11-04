package telegram

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"read-adviser-bot/lib/e"
	"strconv"
)

type Client struct {
	host     string
	basePath string
	client   http.Client
}

const (
	getUpdatesMethod = "getUpdates"
	sendMessageMethod = "sendMessage"
)

func New(host, token string) *Client {
	return &Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

func newBasePath(token string) string {
	return "bot" + token
}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

	// do request
	data, err := c.doRequest(getUpdatesMethod, q)
	if err != nil {
		return nil, err
	}

	var resp UpdatesReponse

	if err := json.Unmarshal(data, &resp); err!= nil {
		return nil, err
	}
    
}

func (c *Client) doRequest(method string, query url.Values) (data []byte, err error) {
	defer func() {
		err = e.WrapIfErr("failed to make request", err)
	}()

	u := url.URL{
		Scheme:   "https",
		Host:     c.host,
		Path:     path.Join(c.basePath, method),
		RawQuery: query.Encode(),
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) SendMessage(chatID string, text string) error {
	q := url.Values{}
    q.Add("chat_id", chatID)
    q.Add("text", text)

    _, err := c.doRequest(sendMessageMethod, q)
    if err!= nil {
        return e.Wrap("failed to send message", err)
    }

	return nil
}

