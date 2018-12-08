package waservice

// WhatsappServiceInterface ...
type WhatsappServiceInterface interface {
	Login()
	SendTextMessage(string, string)
}
