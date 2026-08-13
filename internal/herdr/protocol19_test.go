package herdr

import (
	"context"
	"testing"
)

// The wire gate must accept the targeted Herdr release: 0.8.0 reports protocol 19.
func TestHandshakeAcceptsHerdr080Protocol19(t *testing.T) {
	t.Parallel()
	s := newServer(func(req map[string]any) []byte {
		return reply(req, map[string]any{
			"type": "pong", "version": "0.8.0", "protocol": 19,
			"capabilities": map[string]any{"live_handoff": true, "detached_server_daemon": true},
		})
	})
	c := newTestClient(t, s)
	p, err := c.Handshake(context.Background())
	if err != nil {
		t.Fatalf("herdr 0.8.0 (protocol 19) must be accepted: %v", err)
	}
	if p.Protocol != 19 || p.Version != "0.8.0" {
		t.Fatalf("unexpected pong: %+v", p)
	}
}

// The gate stays exact-match: a superseded protocol must surface as incompatible
// rather than be tolerated, since mismatched schemas corrupt later decodes.
func TestHandshakeRejectsSupersededProtocol17(t *testing.T) {
	t.Parallel()
	s := newServer(func(req map[string]any) []byte {
		return reply(req, map[string]any{"type": "pong", "version": "0.7.5", "protocol": 17})
	})
	c := newTestClient(t, s)
	if _, err := c.Handshake(context.Background()); !IsCode(err, CodeIncompatible) {
		t.Fatalf("protocol 17 must be rejected once 19 is required, got %v", err)
	}
}
