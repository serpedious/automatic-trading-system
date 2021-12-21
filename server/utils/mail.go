package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
)

func Mail() {
	from := "daiaojob@gmail.com"
	password := "2serp5serp2"

	to := []string{
		"daiaorab@gmail.com",
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	t, _ := template.ParseFiles(wd + "/mail_tmp/test_template.html")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: This is a test subject \n%s\n\n", mimeHeaders)))

	t.Execute(&body, struct {
		Name    string
		Message string
	}{
		Name:    "I am developer",
		Message: "This is a test message in a HTML template",
	})

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
