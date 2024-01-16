package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joefazee/mailer"
)

func main() {
	fmt.Println("START!!")


	var Port int
	Host := os.Getenv("Host")
	Port, _ = strconv.Atoi(os.Getenv("Port"))
	Username := os.Getenv("Username")
	Password := os.Getenv("Password")
	Sender := os.Getenv("Sender")

	config := mailer.MailerConfig{
		Host:     Host,
		Port:     Port,
		Username: Username,
		Password: Password,
		Timeout:  5 * time.Second,
		Sender:   Sender,
	}

	to := os.Getenv("Reciever")
	sender := mailer.New(config)

	err := sender.Send(to, "welcome.html", nil)
	if err != nil {
		fmt.Println("error !!")
		log.Fatal(err)
	}

	fmt.Printf("Email sent for %v \n", to)
}