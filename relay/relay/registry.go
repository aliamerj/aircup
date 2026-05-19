package relay

import "sync"

type Registry struct {
	mu       sync.RWMutex
	sessions map[string]*Session
}

func NewRegistry() *Registry {
	return &Registry{
		sessions: make(map[string]*Session),
	}
}

func (r *Registry) Add(nodeID string, s *Session) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.sessions[nodeID] = s
}

func (r *Registry) Remove(nodeID string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.sessions, nodeID)
}
