package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("simple case", func(t *testing.T) {
		err := Copy("testdata/input.txt", "out.txt", 0, 0)
		require.NoError(t, err)
	})

	t.Run("offset exceeds file size", func(t *testing.T) {
		err := Copy("testdata/input.txt", "out.txt", 10000, 0)
		require.ErrorIs(t, err, ErrOffsetExceedsFileSize)
	})

	t.Run("the limit is more than the file size", func(t *testing.T) {
		err := Copy("testdata/input.txt", "out.txt", 0, 100000)
		require.NoError(t, err)
	})
}
