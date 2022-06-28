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

	msg, err := c.Get("<transactionID>")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("Status:", msg.StatusCode)
}
