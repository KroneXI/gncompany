package models

import "sync"

// Metrics содержит данные о посещениях
type Metrics struct {
	TotalVisits int
	mu          sync.Mutex
}

// Increment увеличивает счетчик посещений
func (m *Metrics) Increment() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.TotalVisits++
}

// GetTotalVisits возвращает текущее количество посещений
func (m *Metrics) GetTotalVisits() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.TotalVisits
}
