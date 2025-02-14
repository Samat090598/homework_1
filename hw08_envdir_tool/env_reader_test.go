package main

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	dir, _ := os.Getwd()

	dir = path.Join(dir, "testdata/env")

	t.Run("simple case", func(t *testing.T) {
		environment := Environment{
			"EMPTY": EnvValue{
				Value: "",
			},
			"FOO": EnvValue{
				Value: `   foo
with new line`,
			},
			"HELLO": EnvValue{
				Value: `"hello"`,
			},
			"UNSET": EnvValue{
				NeedRemove: true,
			},
			"BAR": EnvValue{
				Value: `bar`,
			},
		}

		result, _ := ReadDir(dir)

		require.Equal(t, environment, result)
	})
}
