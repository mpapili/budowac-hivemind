// Command hivemind drives high-level NPC AI for Budowac.
package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/mpapili/budowac-hivemind/internal/brain"
	"github.com/mpapili/budowac-hivemind/internal/bus"
	"github.com/mpapili/budowac-hivemind/internal/emit"
	"github.com/mpapili/budowac-hivemind/internal/ingest"
	"github.com/mpapili/budowac-hivemind/internal/proto"
)

func main() {
	port := envOr("PORT", "8092")
	gameID := envOr("GAME_ID", "local-dev")
	natsURL := envOr("NATS_URL", "nats://localhost:4222")

	nc := bus.Connect(natsURL)
	if nc != nil {
		defer nc.Drain()
	}

	reg := brain.DefaultRegistry()
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	var seq atomic.Uint64
	var lastNPC atomic.Int64

	ingest.SubscribeNPCs(nc, gameID, func(report proto.NPCReport) {
		lastNPC.Store(int64(report.Tick))
		for _, n := range report.NPCs {
			b, ok := reg.For(n.Kind)
			if !ok {
				continue
			}
			s := seq.Add(1)
			intent := b.Tick(n, rng, s)
			if intent == nil {
				continue
			}
			if err := emit.Publish(nc, gameID, *intent); err != nil {
				log.Printf("hivemind: publish: %v", err)
			}
		}
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"status":  "ok",
			"service": "budowac-hivemind",
			"gameId":  gameID,
			"lastNpc": lastNPC.Load(),
		})
	})

	srv := &http.Server{Addr: ":" + port, Handler: mux, ReadHeaderTimeout: 5 * time.Second}
	go func() {
		log.Printf("budowac-hivemind listening on http://localhost:%s game=%s", port, gameID)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	log.Printf("budowac-hivemind stopped")
}

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
