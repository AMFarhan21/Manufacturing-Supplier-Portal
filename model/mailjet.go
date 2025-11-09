package model

type Mailjet struct {
	Messages []Message `json:"Messages"`
}

type Message struct {
	From     From   `json:"From"`
	To       []From `json:"To"`
	Subject  string `json:"Subject"`
	TextPart string `json:"TextPart"`
	HTMLPart string `json:"HTMLPart"`
}

type From struct {
	Email string `json:"Email"`
	Name  string `json:"Name"`
}
