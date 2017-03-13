package pkg

import "net/url"

type BotInfo struct {
	URL        *url.URL
	WebhookURL *url.URL
}

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	UserName  string `json:"username,omitempty"`
}

type Chat struct {
	ID                          int    `json:"id"`
	Type                        string `json:"type"`
	Title                       string `json:"title,omitempty"`
	Username                    string `json:"username,omitempty"`
	FirstName                   string `json:"first_name,omitempty"`
	LastName                    string `json:"last_name,omitempty"`
	AllMembersAreAdministrators bool
}

type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url,omitempty"`
	User   User   `json:"user,omitempty"`
}

type Message struct {
	MessageID int             `json:"message_id"`
	From      User            `json:"from,omitempty"`
	Date      int             `json:"date"`
	Chat      Chat            `json:"chat"`
	Text      string          `json: "text,omitempty"`
	Entities  []MessageEntity `json:"entities,omitempty"`
}

type Update struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Response struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type ResponseSentMessage struct {
	OK          bool    `json:"ok"`
	Result      Message `json:"result,omitempty"`
	ErrorCode   int     `json:"error_code,omitempty"`
	Description string  `json:"description,omitempty"`
}