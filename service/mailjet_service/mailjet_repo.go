package mailjet_service

type MailjetRepo interface {
	SendMailjetMessage(senderEmail string, senderName string, receiverEmail string, receiverName string) error
}
