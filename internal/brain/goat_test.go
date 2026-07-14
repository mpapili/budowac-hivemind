package brain

import (
	"math/rand"
	"testing"

	"github.com/mpapili/budowac-hivemind/internal/proto"
)

func TestGoatAssignsWhenIdle(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	intent := Goat{}.Tick(proto.NPCStatus{EntityID: "goat-1", Kind: "goat", X: 0, Z: 0, HasGoal: false}, r, 1)
	if intent == nil {
		t.Fatal("expected intent when idle")
	}
	if intent.EntityID != "goat-1" || intent.Kind != "goat" {
		t.Errorf("bad intent %+v", intent)
	}
	if intent.GoalX < -goalRange || intent.GoalX > goalRange {
		t.Errorf("goalX out of range: %f", intent.GoalX)
	}
}

func TestGoatSilentWhenBusy(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	intent := Goat{}.Tick(proto.NPCStatus{
		EntityID: "goat-1", Kind: "goat",
		X: 0, Z: 0, GoalX: 20, GoalZ: 20, HasGoal: true,
	}, r, 2)
	if intent != nil {
		t.Fatalf("expected nil while far from goal, got %+v", intent)
	}
}

func TestGoatReassignsNearGoal(t *testing.T) {
	r := rand.New(rand.NewSource(2))
	intent := Goat{}.Tick(proto.NPCStatus{
		EntityID: "goat-1", Kind: "goat",
		X: 5, Z: 5, GoalX: 5.2, GoalZ: 5.1, HasGoal: true,
	}, r, 3)
	if intent == nil {
		t.Fatal("expected reassign near goal")
	}
}

func TestRegistryGoat(t *testing.T) {
	reg := DefaultRegistry()
	b, ok := reg.For("goat")
	if !ok || b == nil {
		t.Fatal("goat not registered")
	}
	if _, ok := reg.For("wolf"); ok {
		t.Fatal("unexpected wolf")
	}
}

func TestSurfaceHeightRange(t *testing.T) {
	h := SurfaceHeight(0, 0, 42)
	if h < 6 || h > 20 {
		t.Errorf("h=%d", h)
	}
}

// Client generator.ts surfaceHeight(seed=42) samples — keep parity with budowac-server.
func TestSurfaceHeightMatchesClient(t *testing.T) {
	if got := SurfaceHeight(0, 0, 42); got != 14 {
		t.Errorf("SurfaceHeight(0,0)=%d want 14", got)
	}
	if got := SurfaceHeight(10, 10, 42); got != 12 {
		t.Errorf("SurfaceHeight(10,10)=%d want 12", got)
	}
}
