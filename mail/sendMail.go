package sendMail

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendMail() {
	m := []string{"example@gmail.com"}
	msg := []byte("Hei Hei")

	mail, key, provider := LoadEnv()

	auth := smtp.PlainAuth("", mail, key, provider)
	fmt.Println("Sending mail...")
	err := smtp.SendMail(provider+":587", auth, mail, m, msg)
	if err != nil {
		log.Fatalf("Error sending the mail: %s", err)
	} else {
		fmt.Printf("Mail sent to: %v\n", m)
	}

}

func LoadEnv() (string, string, string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s \n", err)
	}

	key := os.Getenv("MAIL_KEY")
	mail := os.Getenv("MAIL")
	provider := os.Getenv("PROVIDER")

	return mail, key, provider
}
