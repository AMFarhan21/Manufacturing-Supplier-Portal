package xendit

import (
	"Manufacturing-Supplier-Portal/service/xendit_service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type XenditRepository struct {
	xenditApi          string
	xenditUrl          string
	xenditWebhookUrl   string
	successRedirectUrl string
	failureRedirectUrl string
}

func NewXenditRepository(xenditApi, xenditUrl, xenditWebhookUrl, successRedirectUrl, failureRedirectUrl string) *XenditRepository {
	return &XenditRepository{
		xenditApi:          xenditApi,
		xenditUrl:          xenditUrl,
		xenditWebhookUrl:   xenditWebhookUrl,
		successRedirectUrl: successRedirectUrl,
		failureRedirectUrl: failureRedirectUrl,
	}
}

func (r XenditRepository) XenditInvoiceUrl(userId, description, username, email, name, category string, paymentId int, amount float64) (string, error) {

	url := r.xenditUrl
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{
		"external_id": "%d|%s",
		"amount": %.2f,
		"description": "%s",
		"invoice_duration": 3600,
		"customer": {
			"email": "%s"
		},
		"success_redirect_url": "%s",
		"failure_redirect_url": "%s",
		"currency": "IDR",
		"items": [
			{
			"name": "%s",
			"quantity": 1,
			"price": %.2f,
			"category": "%s"
			}
		],
		"metadata": {
			"store_branch": "Makassar",
			"user_id": "%s",
			"payment_id": "%d",
			"username": "%s"
		}
	}      `, paymentId, userId, amount, description, email, r.successRedirectUrl, r.failureRedirectUrl, name, amount, category, userId, paymentId, username))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", r.xenditApi))
	req.Header.Add("Cookie", "__cf_bm=_y6J2eEmO2_wiPddsvXgUQS24DJdIlPIDViHq8aEa4c-1762356798.765628-1.0.1.1-5F1zRs5pVcS07hwmvinbN239JL7gVaEm_IE0ocMvmLg79mWOrcvcuVYPjuaMQLDGI49MIp3ACXcwfnbcgXrH6kN_MYkpd6p7autz.xSS8E9aKC.eqVUKb09MH69j_udx")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	// fmt.Println(string(body))

	var xenditReponse xendit_service.XenditResponse
	err = json.Unmarshal(body, &xenditReponse)
	if err != nil {
		return "", nil
	}

	return xenditReponse.InvoiceURL, nil
}

// func (r XenditRepository) XenditWebhook(paymentId int) (string, error) {
// 	url := r.xenditWebhookUrl
// 	method := "GET"

// 	payload := strings.NewReader(``)

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, url, payload)

// 	if err != nil {
// 		return "", err
// 	}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return "", err
// 	}
// 	// fmt.Println(string(body))

// 	var webHookResponse xendit_service.WebhookResponse
// 	err = json.Unmarshal(body, webHookResponse)
// 	if err != nil {
// 		return "", err
// 	}

// 	var WebHookContentResponse []xendit_service.WebHookContentResponse
// 	for _, data := range webHookResponse.Data {
// 		_ := json.Unmarshal([]byte(data.Content), WebHookContentResponse)
// 	}

// 	var status string
// 	for _, content := range WebHookContentResponse {
// 		if content.ExternalID == string(paymentId) {
// 			status = content.Status
// 		}
// 	}

// 	return status, nil
// }
