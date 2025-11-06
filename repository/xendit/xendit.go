package xendit

import (
	"Manufacturing-Supplier-Portal/service/xendit_service"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
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

func (r XenditRepository) XenditInvoiceUrl(userId, description, username, email, name, category string, rentalId int, amount float64) (string, error) {

	url := r.xenditUrl
	method := "POST"
	externalId := fmt.Sprintf("%v-%s-%d", time.Now(), userId, rentalId)

	fmt.Println("XENDITAPI", r.xenditApi)
	fmt.Println("XENDITURL", r.xenditUrl)
	fmt.Println("SUCCESSREDIRECTURL", r.successRedirectUrl)
	fmt.Println("FAILUREREDIRECTURL", r.failureRedirectUrl)

	payload := strings.NewReader(fmt.Sprintf(`{
		"external_id": "%s",
		"amount": %.2f,
		"description": "%s",
		"invoice_duration": 3600,
		"customer": {
			"username": "%s",
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
			"store_branch": "Makassar"
		}
	}      `, externalId, amount, description, username, email, r.successRedirectUrl, r.failureRedirectUrl, name, amount, category))

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
	fmt.Println(string(body))

	var xenditReponse xendit_service.XenditResponse
	err = json.Unmarshal(body, &xenditReponse)
	if err != nil {
		return "", nil
	}

	return xenditReponse.InvoiceURL, nil
}

func (r XenditRepository) XenditWebhook() {
	url := r.xenditWebhookUrl
	method := "GET"

	payload := strings.NewReader(``)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
