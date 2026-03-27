package client

import "sync"

type clientState struct {
	mu       sync.RWMutex
	apiReady bool
	reqId    int32
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

func (s *clientState) getNextReqId() int32 {
	if !s.isReady() {
		return -1
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	res := s.reqId
	s.reqId++
	return res
}

func (s *clientState) setNextOrderId(id int32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.orderId = id
}

func (s *clientState) getNextOrderId() int32 {
	if !s.isReady() {
		return -1
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	res := s.orderId
	s.orderId++
	return res
}
