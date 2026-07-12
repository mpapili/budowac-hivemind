package bus

import "testing"

func TestSubjects(t *testing.T) {
	if got := ServerNPCs("g1"); got != "budowac.server.g1.npcs" {
		t.Errorf("ServerNPCs = %q", got)
	}
	if got := HivemindIntent("g1"); got != "budowac.hivemind.g1.intent" {
		t.Errorf("HivemindIntent = %q", got)
	}
}
