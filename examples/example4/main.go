package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/madsaune/target365-sdk-go/environments"
	"github.com/madsaune/target365-sdk-go/services/outmessage"
)

func main() {
	token := os.Getenv("STREX_TOKEN")

	client := outmessage.NewClient(string(environments.BaseURLShared), token)

	sendTime := time.Now().Add(time.Minute * 5)

	message := &outmessage.Message{
		Sender:    "Contoso",
		Recipient: "+4799999999",
		Content:   "Hello, World!\n\nThis is an example text message :)",
		SendTime:  &sendTime,
	}

	// Creates a new out-message a.k.a sends it
	resp, err := client.Create(message)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("SMS was created!")
	fmt.Println("ID:", resp.TransactionID)
	fmt.Println("Location:", resp.Location)
}
