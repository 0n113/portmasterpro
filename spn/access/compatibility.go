// Copyright (c) 2026
// SPDX-License-Identifier: GPL-3.0-or-later

// This file preserves the minimal public surface needed while the former
// Safing account backend is intentionally removed. Every function is local,
// side-effect free, and performs no network, token, account, or telemetry I/O.
package access

import "errors"

// ErrNotLoggedIn is retained for source compatibility only. No account login
// exists in portmasterpro; callers must not treat it as a recoverable login flow.
var ErrNotLoggedIn = errors.New("account authentication has been removed from portmasterpro")

// UserRecord is a deliberately empty compatibility record. Account data is not
// stored, loaded, or transmitted by portmasterpro.
type UserRecord struct{}

// InitializeZones is a no-op because account-associated SPN zones are removed.
func InitializeZones() error { return nil }

// GetUser always reports that accounts are unavailable.
func GetUser() (*UserRecord, error) { return nil, ErrNotLoggedIn }

// Login never performs authentication or network I/O.
func Login(_ string, _ string) (*UserRecord, error) { return nil, ErrNotLoggedIn }

// UpdateUser never stores account information.
func UpdateUser(_ *UserRecord) error { return ErrNotLoggedIn }

// GetAuthToken never exposes or fetches an authentication token.
func GetAuthToken() (string, error) { return "", ErrNotLoggedIn }

// UpdateTokens never stores credentials or tokens.
func UpdateTokens(_ ...string) error { return ErrNotLoggedIn }

// GetTokenAmount returns zero because tokens and subscriptions are removed.
func GetTokenAmount() int { return 0 }

// ExpandAndConnectZones is a no-op. It returns ErrNotLoggedIn to make removed
// SPN functionality explicit while preserving a local, offline failure mode.
func ExpandAndConnectZones(_ ...interface{}) error { return ErrNotLoggedIn }
