package main

import (
	"fmt"
	"net/smtp"
)

func main() {

	//targetMail := flag.String("destination", "", "target email")
	//flag.Parse()
	//
	//if *targetMail == "" {
	//	log.Fatal("please specify destination email address")
	//	os.Exit(0)
	//}

	// Sender data.
	from := "marcondol.bot@gmail.com"
	password := "D3ploymentbot"

	// Receiver email address.
	//to := []string{*targetMail}
	to := []string{"azwar.nrst@gmail.com"}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message from marcondol bot.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
