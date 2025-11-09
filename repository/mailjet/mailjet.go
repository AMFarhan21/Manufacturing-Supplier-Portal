package mailjet

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pobyzaarif/goshortcute"
)

type MailJet struct {
	mailjetURL    string
	mailjetAPI    string
	mailjetSECRET string
}

type PayloadSendEmail struct {
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

func NewMailjet(mailjetURL string, mailjetAPI string, mailjetSECRET string) *MailJet {
	return &MailJet{
		mailjetURL:    mailjetURL,
		mailjetAPI:    mailjetAPI,
		mailjetSECRET: mailjetSECRET,
	}
}

func (r MailJet) SendMailjetMessage(senderEmail string, senderName string, receiverEmail string, receiverName string) error {
	url := r.mailjetURL
	method := "POST"

	message := Message{
		From: From{
			Email: senderEmail,
			Name:  senderName,
		},
		To: []From{
			{
				Email: senderEmail,
				Name:  senderName,
			},
		},
		Subject:  "Validate your email",
		TextPart: "TESTING",
		HTMLPart: fmt.Sprintf(`<h3>Validate your email using this link <a href=\"https://manufacturing-supplier-portal.onrender.com\api\ValidateEmailAddress\">Validate your email!</a><h3><br/> Welcome %s`, receiverName),
	}

	payload := []PayloadSendEmail{{
		[]Message{message},
	}}

	payloadByte, _ := json.Marshal(payload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(string(payloadByte)))

	if err != nil {
		return err
	}

	authorization := goshortcute.StringtoBase64Encode(r.mailjetAPI + ":" + r.mailjetSECRET)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+authorization)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
