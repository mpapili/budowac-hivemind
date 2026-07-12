package emit

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mpapili/budowac-hivemind/internal/bus"
	"github.com/mpapili/budowac-hivemind/internal/proto"
	"github.com/nats-io/nats.go"
)

// Publish sends an AIIntent on the hivemind subject.
func Publish(nc *nats.Conn, gameID string, intent proto.AIIntent) error {
	if nc == nil {
		return fmt.Errorf("nats not connected")
	}
	b, err := json.Marshal(intent)
	if err != nil {
		return err
	}
	subj := bus.HivemindIntent(gameID)
	if err := nc.Publish(subj, b); err != nil {
		return err
	}
	log.Printf("emit: %s entity=%s goal=(%.1f,%.1f)", subj, intent.EntityID, intent.GoalX, intent.GoalZ)
	return nil
}
