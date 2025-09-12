package handlers

import (
	"context"
	"time"

	"github.com/mk6i/retro-aim-server/state"
)

// BuddyBroadcaster broadcasts buddy presence updates (local copy to avoid import cycle)
type BuddyBroadcaster interface {
	BroadcastBuddyArrived(ctx context.Context, sess *state.Session) error
	BroadcastBuddyDeparted(ctx context.Context, sess *state.Session) error
}

// TokenStore manages authentication tokens (local copy to avoid import cycle)
type TokenStore interface {
	StoreToken(token string, screenName state.IdentScreenName, expiresAt time.Time) error
	ValidateToken(token string) (state.IdentScreenName, error)
	DeleteToken(token string) error
}
