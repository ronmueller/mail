package main

import (
	"encoding/base64"
	"fmt"
	"net/smtp"
	"time"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"",
		"",
		"",
		"localhost",
	)

	// header information
	header := make(map[string]string)
	header["Date"] = time.Now().Format(time.RFC822Z)
	header["From"] = "foxy@unity.ch"
	header["To"] = "ronald.mueller@20minuten.ch,ronmueller@swissonline.ch"
	header["Subject"] = "test mail"
	header["MIME-Version"] = "1.0"
	header["Content-Type"] = "text/plain; charset=\"utf-8\""
	header["Content-Transfer-Encoding"] = "base64"

	// encoded message
	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + base64.StdEncoding.EncodeToString([]byte("ein test email"))

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"localhost:25",
		auth,
		"foxy@unity.ch",
		[]string{"ronald.mueller@20minuten.ch", "ronmueller@swissonline.ch"},
		[]byte(message),
	)

	if err != nil {
		fmt.Println(err)
	}
}
