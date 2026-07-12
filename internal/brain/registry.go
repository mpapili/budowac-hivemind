package brain

// Registry maps entity kind → Behavior. Expandable: Register("wolf", …).
type Registry struct {
	byKind map[string]Behavior
}

func NewRegistry() *Registry {
	return &Registry{byKind: make(map[string]Behavior)}
}

func (r *Registry) Register(kind string, b Behavior) {
	if kind == "" || b == nil {
		return
	}
	r.byKind[kind] = b
}

func (r *Registry) For(kind string) (Behavior, bool) {
	b, ok := r.byKind[kind]
	return b, ok
}

// DefaultRegistry installs all known brains (goat, …).
func DefaultRegistry() *Registry {
	r := NewRegistry()
	r.Register("goat", Goat{})
	return r
}
