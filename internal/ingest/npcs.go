package ingest

import (
	"encoding/json"
	"log"

	"github.com/mpapili/budowac-hivemind/internal/bus"
	"github.com/mpapili/budowac-hivemind/internal/proto"
	"github.com/nats-io/nats.go"
)

// SubscribeNPCs invokes onReport for each server NPCReport.
func SubscribeNPCs(nc *nats.Conn, gameID string, onReport func(proto.NPCReport)) {
	if nc == nil {
		return
	}
	_, err := nc.Subscribe(bus.ServerNPCs(gameID), func(msg *nats.Msg) {
		var r proto.NPCReport
		if err := json.Unmarshal(msg.Data, &r); err != nil {
			log.Printf("ingest: bad npcs: %v", err)
			return
		}
		if onReport != nil {
			onReport(r)
		}
	})
	if err != nil {
		log.Printf("ingest: sub: %v", err)
	}
}
