//This package is just for experimenting with basic channels usage.
package main

import (
	"fmt"
	"strings"
)

type Message struct {
	To      []string
	From    string
	Content string
}

type FailedMessage struct {
	ErrorMessage    string
	OriginalMessage Message
}

func loops() {
	phrase := "These are the times that try men's soul\n"

	words := strings.Split(phrase, " ")
	ch := make(chan string, len(words))
	for _, word := range words {
		ch <- word
	}

	close(ch)

	for msg := range ch {
		fmt.Print(msg + " ")
	}
}

func main() {
	loops()
	msgCh := make(chan Message, 1)
	errCh := make(chan FailedMessage, 1)

	msg := Message{
		To:      []string{"abc@abc.com"},
		From:    "xyz@xyz.me",
		Content: "vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv",
	}

	failedMessage := FailedMessage{
		ErrorMessage:    "Message Error",
		OriginalMessage: Message{},
	}

	msgCh <- msg
	errCh <- failedMessage

	select {
	case recievedMsg := <-msgCh:
		fmt.Println(recievedMsg)
	case recievedErr := <-errCh:
		fmt.Println(recievedErr)
	default:
		fmt.Println("No Message Recieved")
	}
}
