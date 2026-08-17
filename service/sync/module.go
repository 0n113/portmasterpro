// Package sync has been stubbed out.
// The original sync module synchronized settings and profiles with the
// Safing account backend (telemetry / cloud sync). This functionality
// has been intentionally removed from portmasterpro.
package sync

import (
	"github.com/safing/portmaster/service/mgr"
)

// Sync is a no-op stub that satisfies the module interface without
// making any outbound connections or syncing data to external servers.
type Sync struct {
	m *mgr.Manager
}

// New returns a no-op Sync stub.
func New(_ interface{}) (*Sync, error) {
	return &Sync{}, nil
}

// Manager returns nil — no manager needed for a no-op module.
func (s *Sync) Manager() *mgr.Manager {
	return s.m
}

// Start is a no-op.
func (s *Sync) Start() error {
	return nil
}

// Stop is a no-op.
func (s *Sync) Stop() error {
	return nil
}
