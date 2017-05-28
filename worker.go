package messagebus

import (
	"runtime"
	"time"
)

type (
	Worker struct {
		SleepDuration time.Duration
		Listeners     []MessageListener
	}
)

func (w *Worker) Run(r MessageReader) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	for {
		messages := r.Read()
		if len(messages) < 1 {
			time.Sleep(w.SleepDuration)
			continue
		}

		for i := 0; i < len(w.Listeners); i++ {
			go w.Listeners[i].Process(&messages)
		}

		r.AckMessages(&messages)
	}
}
