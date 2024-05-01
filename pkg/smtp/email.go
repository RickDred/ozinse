package smtp

import (
	"net/smtp"
)

func SendEmail(Identity, username, password, host, addr, from, to, subject, body string) error {
	auth := smtp.PlainAuth(Identity, username, password, host)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	err := smtp.SendMail(addr, auth, from, []string{to}, msg)
	if err != nil {
		return err
	}

	return nil
}
