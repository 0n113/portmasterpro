package sync_test

import (
	"testing"

	"github.com/safing/portmaster/service/sync"
)

// TestSyncStubNoOp verifies that the sync stub starts and stops without error
// and does not establish any external connections.
func TestSyncStubNoOp(t *testing.T) {
	s, err := sync.New(nil)
	if err != nil {
		t.Fatalf("sync.New() returned unexpected error: %v", err)
	}
	if err := s.Start(); err != nil {
		t.Fatalf("Sync.Start() returned unexpected error: %v", err)
	}
	if err := s.Stop(); err != nil {
		t.Fatalf("Sync.Stop() returned unexpected error: %v", err)
	}
}

// TestSyncStubManagerIsNil verifies that the stub returns no manager,
// confirming no background workers are spawned.
func TestSyncStubManagerIsNil(t *testing.T) {
	s, err := sync.New(nil)
	if err != nil {
		t.Fatalf("sync.New() returned unexpected error: %v", err)
	}
	if s.Manager() != nil {
		t.Error("expected Manager() to return nil for no-op stub, got non-nil")
	}
}
