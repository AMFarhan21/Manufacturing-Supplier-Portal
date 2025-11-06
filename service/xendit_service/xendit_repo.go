package xendit_service

type XenditRepo interface {
	XenditInvoiceUrl(userId, description, username, email, name, category string, paymentId int, amount float64) (string, error)
	// XenditWebhook()
}
