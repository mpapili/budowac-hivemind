package bus

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func Connect(url string) *nats.Conn {
	if url == "" {
		url = nats.DefaultURL
	}
	var nc *nats.Conn
	var err error
	for attempt := 1; attempt <= 30; attempt++ {
		nc, err = nats.Connect(url,
			nats.Name("budowac-hivemind"),
			nats.MaxReconnects(-1),
			nats.ReconnectWait(2*time.Second),
		)
		if err == nil {
			log.Printf("NATS connected (%s)", url)
			return nc
		}
		log.Printf("NATS connect attempt %d: %v", attempt, err)
		time.Sleep(time.Second)
	}
	log.Printf("NATS unavailable after retries; continuing without bus")
	return nil
}
