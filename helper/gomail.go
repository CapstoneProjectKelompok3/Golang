package helper

import (
	"errors"
	"fmt"

	"gopkg.in/gomail.v2"
)

type MessageGomail struct {
	EmailReceiver string
	Sucject       string
	Content       string
}

const (
	email_from = "agoverment55@gmail.com"
	pass_app   = "xsxd htot sdby xikp"
	actionURL  = "https://belanjalagiyuk.shop/emergencies/action"
)

func SendGomailMessage(input MessageGomail) (string, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", email_from)
	m.SetHeader("To", input.EmailReceiver)
	m.SetHeader("Subject", input.Sucject)
	htmlBody := HTMLImergency()
	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer("smtp.gmail.com", 587, email_from, pass_app)

	if err := d.DialAndSend(m); err != nil {
		return "", errors.New("gagal mengirim email")
	} else {
		return "Email terkirim", nil
	}
}

func HTMLImergency() string {
	htmlBody := fmt.Sprintf(`
        <html>
            <body>
                <h1>Tawaran Anda</h1>
                <p>Ini adalah tawaran dari kami.</p>
                <a href="%s?accept=true">Terima</a>
                <a href="%s?accept=false">Tolak</a>
            </body>
        </html>
    `, actionURL, actionURL)
	return htmlBody
}

func SendResponseEmail(emailAddress, responseMessage string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", email_from)
	m.SetHeader("To", emailAddress)
	m.SetHeader("Subject", "Action Response")
	m.SetBody("text/plain", responseMessage)

	d := gomail.NewDialer("smtp.gmail.com", 587, email_from, pass_app)

	if err := d.DialAndSend(m); err != nil {
		return errors.New("send response failed")
	}
	return nil
}
