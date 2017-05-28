package messagebus

type (
	Message struct {
		Body string
	}

	MessageReader interface {
		Read() []Message
		AckMessages(messages *[]Message)
	}

	MessageListener interface {
		Process(messages *[]Message)
	}
)
