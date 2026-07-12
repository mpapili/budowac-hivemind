package proto

// Wire types duplicated per repo (see budowac-space duplication policy).

type AIIntent struct {
	EntityID string  `json:"entityId"`
	Kind     string  `json:"kind"`
	GoalX    float32 `json:"goalX"`
	GoalY    float32 `json:"goalY"`
	GoalZ    float32 `json:"goalZ"`
	Seq      uint64  `json:"seq"`
}

type NPCStatus struct {
	EntityID string  `json:"entityId"`
	Kind     string  `json:"kind"`
	X        float32 `json:"x"`
	Y        float32 `json:"y"`
	Z        float32 `json:"z"`
	GoalX    float32 `json:"goalX,omitempty"`
	GoalY    float32 `json:"goalY,omitempty"`
	GoalZ    float32 `json:"goalZ,omitempty"`
	HasGoal  bool    `json:"hasGoal"`
}

type NPCReport struct {
	Tick uint64      `json:"tick"`
	NPCs []NPCStatus `json:"npcs"`
}
