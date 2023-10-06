package helper

import (
	"errors"
	"fmt"

	"gopkg.in/gomail.v2"
)

type MessageGomailE struct {
	EmailReceiver string
	Sucject       string
	Content       string
	Name          string
	Email 		  string
}


const (
	email_fromE = "agoverment55@gmail.com"
	pass_appE   = "xsxd htot sdby xikp"
	actionURLE  = "https://belanjalagiyuk.shop/emergencies/action"
)

func SendGomailMessageE(input MessageGomailE) (string, error) {
	m := gomail.NewMessage()
	m.SetHeader("From", email_fromE)
	m.SetHeader("To", input.EmailReceiver)
	m.SetHeader("Subject", input.Sucject)
	htmlBody := HTMLImergencyE(input.Content,input.Name)
	m.SetBody("text/html", htmlBody)

	d := gomail.NewDialer("smtp.gmail.com", 587, email_fromE, pass_appE)

	if err := d.DialAndSend(m); err != nil {
		return "", errors.New("gagal mengirim email")
	} else {
		return "Email terkirim", nil
	}
}

func HTMLImergencyE(content string,name string) string {

	htmlBody := fmt.Sprintf(`
	<html>
	<head>
	<style>
		body {
			font-family: Arial, sans-serif;
			background-color: #f2f2f2;
		}

		.container {
			max-width: 600px;
			margin: 0 auto;
			padding: 20px;
			background-color: #ffffff;
			box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
			text-align: center;
		}

		h1 {
			color: #333333;
		}

		img {
			margin-bottom: 10px;
		}

		p {
			color: #666666;
		}

		button {
			margin: auto;
			margin-bottom: 10px;
			background-color: #007bff;
			color: #ffffff;
			padding: 10px 20px;
			border: none;
			border-radius: 5px;
			text-decoration: none;
			display: inline-block;
			cursor: pointer;
		}

		button a {
			color: white;
			text-decoration: none;
		}

		b {
			font-weight: bold;
		}
  </style>
	</head>
	<body>
		<div class="container">
			<h1>Ada Kasus Baru</h1>

			<img src="https://cdn.discordapp.com/attachments/1155851056740311081/1156266029429837885/logo_ecci.png?ex=651c4127&is=651aefa7&hm=59808162e351fa85b2d3d23d7b06e79ed9069fc617b2691d1b3eb2c759698d6d&">

			<p>
			User dengan nama %s telah melapor tentang: <br>
			%s
			</p>

			<p>Diharapkan admin dapat menangani kasus ini</p>
		</div>
	</body>
</html>
    `,name,content)

	return htmlBody
}

func SendResponseEmailE(emailAddress, responseMessage string) error {
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
