package mailjet

import (
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

func (r MailJet) SendMailjetMessage(senderEmail string, senderName string, receiverEmail string, receiverName string, signedToken string) error {
	url := r.mailjetURL
	method := "POST"
	payload := strings.NewReader(fmt.Sprintf(`{
    "Messages": [
			{
				"From": {
					"Email": "%s",
					"Name": "%s"
				},
				"To": [
					{
						"Email": "%s",
						"Name": "%s"
					}
				],
				"Subject": "Validate your email!",
				"TextPart": "Validate your email https://youtube.com",
				"HTMLPart": "<h3>Click this to <a href=\"https://manufacturing-supplier-portal.onrender.com/api/auth/validateemailaddress?token=%v\">Validate</a> <br/> Welcome %s"
			}
		]
	}`, senderEmail, senderName, receiverEmail, receiverName, signedToken, receiverName))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

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
