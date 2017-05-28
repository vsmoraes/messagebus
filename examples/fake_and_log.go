package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/vsmoraes/messagebus"
)

type (
	FakeReader  struct{}
	LogListener struct{}
)

func (r *FakeReader) Read() []messagebus.Message {
	c := rand.Intn(3)
	var messages []messagebus.Message

	for i := 0; i < c; i++ {
		message := &messagebus.Message{Body: "foo"}
		messages = append(messages, *message)
	}

	return messages
}

func (r *FakeReader) AckMessages(messages *[]messagebus.Message) {
	c := len(*messages)
	log.Println(strconv.Itoa(c) + " messages acknowledged")
}

func (l *LogListener) Process(messages *[]messagebus.Message) {
	c := len(*messages)
	time.Sleep(time.Second * 5)
	log.Println(strconv.Itoa(c) + " messages processed")
}

func main() {
	var listeners []messagebus.MessageListener
	listeners = append(listeners, &LogListener{})
	reader := &FakeReader{}

	bus := &messagebus.Worker{
		Listeners:     listeners,
		SleepDuration: time.Second * 1,
	}

	bus.Run(reader)
}
