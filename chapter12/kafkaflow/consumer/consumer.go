package main

import (
	sarama "github.com/IBM/sarama"
	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter12/internal/kafkaflow"
	flow "github.com/trustmaster/goflow"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9094"}, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("example", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer partitionConsumer.Close()

	net := kafkaflow.NewUpperApp()

	in := make(chan string)
	err = net.SetInPort("In", in)
	if err != nil {
		panic(err)
	}

	wait := flow.Run(net)
	defer func() {
		close(in)
		<-wait
	}()

	for {
		msg := <-partitionConsumer.Messages()
		in <- string(msg.Value)
	}

}
