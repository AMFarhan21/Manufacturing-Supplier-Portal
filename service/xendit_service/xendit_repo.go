package xendit_service

type XenditRepo interface {
	XenditInvoiceUrl(userId, description, username, email, name, category string, rentalId int, amount float64) (string, error)
	XenditWebhook()
}
