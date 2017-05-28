package messagebus

type (
	Message struct {
		Body string
	}

	MessageReader interface {
		Read() []Message
	}

	MessageListener interface {
		Process(messages *[]Message)
	}
)
