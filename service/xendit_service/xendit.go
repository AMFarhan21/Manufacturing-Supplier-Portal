package xendit_service

import "time"

type XenditResponse struct {
	ID                        string                  `json:"id"`
	ExternalID                string                  `json:"external_id"`
	UserID                    string                  `json:"user_id"`
	Status                    string                  `json:"status"`
	MerchantName              string                  `json:"merchant_name"`
	MerchantProfilePictureURL string                  `json:"merchant_profile_picture_url"`
	Amount                    int64                   `json:"amount"`
	Description               string                  `json:"description"`
	ExpiryDate                time.Time               `json:"expiry_date"`
	InvoiceURL                string                  `json:"invoice_url"`
	AvailableBanks            []AvailableBank         `json:"available_banks"`
	AvailableRetailOutlets    []AvailableRetailOutlet `json:"available_retail_outlets"`
	AvailableEwallets         []AvailableEwallet      `json:"available_ewallets"`
	AvailableQrCodes          []AvailableQrCode       `json:"available_qr_codes"`
	AvailableDirectDebits     []AvailableDirectDebit  `json:"available_direct_debits"`
	AvailablePaylaters        []AvailablePaylater     `json:"available_paylaters"`
	ShouldExcludeCreditCard   bool                    `json:"should_exclude_credit_card"`
	ShouldSendEmail           bool                    `json:"should_send_email"`
	SuccessRedirectURL        string                  `json:"success_redirect_url"`
	FailureRedirectURL        string                  `json:"failure_redirect_url"`
	Created                   time.Time               `json:"created"`
	Updated                   time.Time               `json:"updated"`
	Currency                  string                  `json:"currency"`
	Items                     []Item                  `json:"items"`
	Customer                  Customer                `json:"customer"`
	Metadata                  Metadata                `json:"metadata"`
}

type AvailableBank struct {
	BankCode          string `json:"bank_code"`
	CollectionType    string `json:"collection_type"`
	TransferAmount    int64  `json:"transfer_amount"`
	BankBranch        string `json:"bank_branch"`
	AccountHolderName string `json:"account_holder_name"`
	IdentityAmount    int64  `json:"identity_amount"`
}

type AvailableDirectDebit struct {
	DirectDebitType string `json:"direct_debit_type"`
}

type AvailableEwallet struct {
	EwalletType string `json:"ewallet_type"`
}

type AvailablePaylater struct {
	PaylaterType string `json:"paylater_type"`
}

type AvailableQrCode struct {
	QrCodeType string `json:"qr_code_type"`
}

type AvailableRetailOutlet struct {
	RetailOutletName string `json:"retail_outlet_name"`
}

type Customer struct {
	Email string `json:"email"`
}

type Item struct {
	Name     string `json:"name"`
	Quantity int64  `json:"quantity"`
	Price    int64  `json:"price"`
	Category string `json:"category"`
}

type Metadata struct {
	StoreBranch string `json:"store_branch"`
}

type WebhookResponse struct {
	Data []struct {
		UUID        string      `json:"uuid"`
		Type        string      `json:"type"`
		TokenID     string      `json:"token_id"`
		TeamID      interface{} `json:"team_id"`
		IP          string      `json:"ip"`
		Country     string      `json:"country"`
		CountryCode string      `json:"country_code"`
		Region      string      `json:"region"`
		City        string      `json:"city"`
		Hostname    string      `json:"hostname"`
		Method      string      `json:"method"`
		UserAgent   string      `json:"user_agent"`
		Content     string      `json:"content"`
		Query       interface{} `json:"query"`
		Headers     struct {
			AcceptLanguage          []string `json:"accept-language"`
			AcceptEncoding          []string `json:"accept-encoding"`
			Referer                 []string `json:"referer"`
			SecFetchDest            []string `json:"sec-fetch-dest"`
			SecFetchUser            []string `json:"sec-fetch-user"`
			SecFetchMode            []string `json:"sec-fetch-mode"`
			SecFetchSite            []string `json:"sec-fetch-site"`
			Accept                  []string `json:"accept"`
			UserAgent               []string `json:"user-agent"`
			UpgradeInsecureRequests []string `json:"upgrade-insecure-requests"`
			SecChUaPlatform         []string `json:"sec-ch-ua-platform"`
			SecChUaMobile           []string `json:"sec-ch-ua-mobile"`
			SecChUa                 []string `json:"sec-ch-ua"`
			Host                    []string `json:"host"`
		} `json:"headers"`
		URL                string        `json:"url"`
		Size               int           `json:"size"`
		Files              []interface{} `json:"files"`
		CreatedAt          string        `json:"created_at"`
		UpdatedAt          string        `json:"updated_at"`
		Sorting            int64         `json:"sorting"`
		CustomActionOutput []interface{} `json:"custom_action_output"`
		CustomActionErrors []interface{} `json:"custom_action_errors"`
		Time               float64       `json:"time"`
	} `json:"data"`
	Total       int  `json:"total"`
	PerPage     int  `json:"per_page"`
	CurrentPage int  `json:"current_page"`
	IsLastPage  bool `json:"is_last_page"`
	From        int  `json:"from"`
	To          int  `json:"to"`
}

type WebHookContentResponse struct {
	ID                 string                       `json:"id"`
	ExternalID         string                       `json:"external_id"`
	UserID             string                       `json:"user_id"`
	PaymentMethod      string                       `json:"payment_method"`
	Status             string                       `json:"status"`
	MerchantName       string                       `json:"merchant_name"`
	Amount             int64                        `json:"amount"`
	PaidAmount         int64                        `json:"paid_amount"`
	BankCode           string                       `json:"bank_code"`
	PaidAt             time.Time                    `json:"paid_at"`
	Description        string                       `json:"description"`
	IsHigh             bool                         `json:"is_high"`
	SuccessRedirectURL string                       `json:"success_redirect_url"`
	FailureRedirectURL string                       `json:"failure_redirect_url"`
	Created            time.Time                    `json:"created"`
	Updated            time.Time                    `json:"updated"`
	Currency           string                       `json:"currency"`
	PaymentChannel     string                       `json:"payment_channel"`
	PaymentDestination string                       `json:"payment_destination"`
	Items              []WebHookContentResponseItem `json:"items"`
	PaymentID          string                       `json:"payment_id"`
}

type WebHookContentResponseItem struct {
	Name     string `json:"name"`
	Quantity int64  `json:"quantity"`
	Price    int64  `json:"price"`
	Category string `json:"category"`
}
