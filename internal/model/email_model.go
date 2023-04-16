package model

type MailPayload struct {
	From     string
	To       string
	Subject  string
	Template string
	Data     map[string]interface{}
}
