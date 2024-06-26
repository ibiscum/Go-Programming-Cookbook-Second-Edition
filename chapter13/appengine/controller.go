package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/datastore"
)

// Controller holds our storage and other
// state
type Controller struct {
	store *datastore.Client
}

func (c *Controller) handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	ctx := context.Background()

	// store the new message
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	if message := r.FormValue("message"); message != "" {
		if err := c.storeMessage(ctx, message); err != nil {
			log.Printf("could not store message: %v", err)
			http.Error(w, "could not store message", http.StatusInternalServerError)
			return
		}
	}

	// get the current messages and display them
	messages, err := c.queryMessages(ctx, 10)
	if err != nil {
		log.Printf("could not get messages: %v", err)
		http.Error(w, "could not get messages", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Messages:")
	for _, message := range messages {
		fmt.Fprintln(w, message.Message)
	}
}
