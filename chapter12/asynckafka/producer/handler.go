package main

import (
	"net/http"

	sarama "github.com/IBM/sarama"
)

// KafkaController allows us to attach a producer
// to our handlers
type KafkaController struct {
	producer sarama.AsyncProducer
}

// Handler grabs a message from a GET parama and
// send it to the kafka queue asynchronously
func (c *KafkaController) Handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg := r.FormValue("msg")
	if msg == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("msg must be set"))
		if err != nil {
			panic(err)
		}
		return
	}
	c.producer.Input() <- &sarama.ProducerMessage{Topic: "example", Key: nil, Value: sarama.StringEncoder(msg)}
	w.WriteHeader(http.StatusOK)
}
