package sentinel

import (
	"errors"
)

// Sentinel is a client for the Sentinel service.
type Sentinel struct {
	url     string
	timeout int
	token   string
}

var (
	ErrInitializationTimeout = errors.New("timeout waiting for Sentinel client initialization")
	ErrInitializationFailed  = errors.New("")
	ErrClientNotInitialized  = errors.New("")
)

// New creates a new Sentinel client.
func New(options ...Option) (*Sentinel, error) {
	s := &Sentinel{
		getEnv("SENTINEL_ADDR", "http://localhost:8080"),
		10,
		getEnv("SENTINEL_TOKEN", ""),
	}
	for _, option := range options {
		option(s)
	}

	return s, nil
}

type Option func(*Sentinel)

// WithTimeout sets the timeout for the Sentinel client.
func WithTimeout(timeout int) Option {
	return func(s *Sentinel) {
		s.timeout = timeout
	}
}

// WithServer sets the server address for the Sentinel client.
func WithServer(server string) Option {
	return func(s *Sentinel) {
		s.url = server
	}
}

// WithToken sets the token for the Sentinel client.
func WithToken(token string) Option {
	return func(s *Sentinel) {
	}
}

// Approve approves the activity within a workflow
// Reject rejects the activity within a workflow
// Register registers the activity within a workflow
// PendingApprovals returns the pending approvals
// ApprovalRequests returns the approval requests

func (s *Sentinel) Approve(key string) (string, error) {
	return "", nil
}
