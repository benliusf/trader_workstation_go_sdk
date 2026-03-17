package client

import "sync"

type clientState struct {
	mu       sync.RWMutex
	apiReady bool
	orderId  int32
}

func (s *clientState) setReady() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.apiReady = true
}

func (s *clientState) isReady() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.apiReady
}

func (s *clientState) setNextValidId(id int32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.orderId = id
}

func (s *clientState) getNextValidId() int32 {
	s.mu.Lock()
	defer s.mu.Unlock()
	res := s.orderId
	s.orderId += 1
	return res
}
