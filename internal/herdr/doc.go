// Package herdr is the single typed owner of the Herdr socket wire protocol
// (protocol 19, schema version 1, Herdr 0.8.0).
//
// It speaks newline-delimited JSON over a Unix domain socket. Normal requests
// use one connection each with a bounded read, a string request id, result-type
// validation, and a default five-second deadline. A separate long-lived
// [Subscriber] keeps an events.subscribe connection open purely to wake the
// state engine; snapshots remain the source of truth.
//
// The package never forwards arbitrary browser-supplied method names or raw
// params: every request is constructed from a typed method here. Every mutation
// carries explicit resource ids and never relies on Herdr UI focus.
//
// All external time and I/O are injected through [Clock] and [Dialer] so the
// package is fully deterministic under test.
package herdr

// Protocol is the Herdr wire protocol version this package targets.
const Protocol = 19

// MinHerdrVersion is the oldest Herdr release whose APIs this package uses.
const MinHerdrVersion = "0.8.0"
