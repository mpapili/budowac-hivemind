package brain

import (
	"math/rand"

	"github.com/mpapili/budowac-hivemind/internal/proto"
)

// Behavior is an expandable per-kind high-level AI.
// Return nil when no new intent is required.
type Behavior interface {
	Tick(npc proto.NPCStatus, r *rand.Rand, seq uint64) *proto.AIIntent
}
