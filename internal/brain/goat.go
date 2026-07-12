package brain

import (
	"math"
	"math/rand"

	"github.com/mpapili/budowac-hivemind/internal/proto"
)

const (
	worldSeed  int64   = 42
	goalRange  float32 = 40
	arriveDist float32 = 1.0
)

// Goat picks a random surface point when idle or near its current goal.
type Goat struct{}

func (Goat) Tick(npc proto.NPCStatus, r *rand.Rand, seq uint64) *proto.AIIntent {
	if r == nil {
		return nil
	}
	need := !npc.HasGoal
	if npc.HasGoal {
		dx := npc.GoalX - npc.X
		dz := npc.GoalZ - npc.Z
		if float32(math.Hypot(float64(dx), float64(dz))) < arriveDist {
			need = true
		}
	}
	if !need {
		return nil
	}
	gx := (r.Float32()*2 - 1) * goalRange
	gz := (r.Float32()*2 - 1) * goalRange
	gy := float32(SurfaceHeight(int(math.Floor(float64(gx))), int(math.Floor(float64(gz))), worldSeed)) + 0.01
	return &proto.AIIntent{
		EntityID: npc.EntityID,
		Kind:     "goat",
		GoalX:    gx,
		GoalY:    gy,
		GoalZ:    gz,
		Seq:      seq,
	}
}
