package mail

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

// Function to send mail asynchronously
func SendMail(to, subject, body, htmlBody string) error {
	sender := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	host := os.Getenv("SMTP_SERVER")
	portStr := os.Getenv("SMTP_PORT")

	// Convert port string to integer
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)
	m.AddAlternative("text/html", htmlBody)

	d := gomail.NewDialer(host, port, sender, password)

	// Send the email asynchronously
	go func() {
		if err := d.DialAndSend(m); err != nil {
			fmt.Println("Failed to send email:", err)
		} else {
			fmt.Println("Email sent successfully!")
		}
	}()

	return nil
}

func MailPassword(to, password string) error {
    subject := "Your New Password"
    body := fmt.Sprintf("Your new password is: %s", password)
    htmlBody := GeneratePasswordHTML(to, password)

    return SendMail(to, subject, body, htmlBody)
}
