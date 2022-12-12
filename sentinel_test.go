package sentinel

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func test_sentinel_new(t *testing.T) {
	s, err := New()
	require.NoError(t, err)
	s.Register()
	s.Wait()
	s.Approve()
}
