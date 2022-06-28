package main

import (
	"fmt"
	"log"
	"os"

	"github.com/madsaune/target365-sdk-go/client"
	"github.com/madsaune/target365-sdk-go/services/outmessage"
)

func main() {
	token := os.Getenv("STREX_TOKEN")

	c := outmessage.NewClient(client.BaseURLShared, token)

	messages := []outmessage.Message{
		{
			Sender:    "Contoso",
			Recipient: "+4799999997",
			Content:   "Hello, World!\n\nThis is an example text message :)",
		},
		{
			Sender:    "Contoso",
			Recipient: "+4799999998",
			Content:   "Hello, World!\n\nThis is an example text message :)",
		},
		{
			Sender:    "Contoso",
			Recipient: "+4799999999",
			Content:   "Hello, World!\n\nThis is an example text message :)",
		},
	}

	// Creates a new out-message a.k.a sends it
	resp, err := c.CreateBatch(messages)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, v := range resp {
		fmt.Println("ID:", v.TransactionID)
		fmt.Println("Location:", v.Location)
		fmt.Println("---")
	}
}
