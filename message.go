package messagebus

type (
	Message struct {
		Body       string
		RawMessage interface{}
	}

	MessageReader interface {
		Read() []Message
		AckMessages(messages *[]Message)
	}

	MessageListener interface {
		Process(messages *[]Message)
	}
)
