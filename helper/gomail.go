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
	htmlBody := HTMLImergency(input.Content)
	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer("smtp.gmail.com", 587, email_from, pass_app)

	if err := d.DialAndSend(m); err != nil {
		return "", errors.New("gagal mengirim email")
	} else {
		return "Email terkirim", nil
	}
}

func HTMLImergency(content string) string {

	htmlBody := fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<style>
			body {
				font-family: Arial, sans-serif;
				background-color: #f5f5f5;
				margin: 0;
				padding: 0;
			}
			h1 {
				color: #333;
			}
			p {
				color: #555;
			}
			h3 {
				color: #333;
			}
			table {
				margin-top: 20px;
			}
			td {
				padding: 10px;
			}
    
		 a.accept-button {
       		 background-color: #008000;
       		 color: #fff; 
       		 text-decoration: none;
       		 padding: 10px 20px;
       		 border-radius: 5px;
    	}
    
    	a.reject-button {
    	    background-color: #FF0000;
    	    color: #fff; /* Warna teks tombol Tolak */
    	    text-decoration: none;
    	    padding: 10px 20px;
    	    border-radius: 5px;
    	}

    
    	a.accept-button,
    	a.reject-button {
    	    color: #fff;
    	}

		</style>
	</head>
	<body>
		<h1>Ada keadaan darurat</h1>
		<p>%s</p>
		<h3>Konfirmasi Pesan Anda Untuk Menghandel keadaan</h3>
		<p>Silakan konfirmasi pesan ini dengan mengklik salah satu tombol di bawah ini:</p>
		<table cellspacing="10" cellpadding="0" border="0">
			<tr>
				<td style="padding: 10px;">
					<a href="%s?accept=true" class="accept-button">Terima</a>
				</td>
				<td style="padding: 10px;">
					<a href="%s?accept=false" class="reject-button">Tolak</a>
				</td>
			</tr>
		</table>
	</body>
	</html>
	
    `,content, actionURL, actionURL)

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
