package util_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thatmattlove/addr/internal/util"
)

func Test_PathExists(t *testing.T) {
	td := t.TempDir()
	tf, err := os.CreateTemp(td, "*")
	assert.NoError(t, err)
	defer tf.Close()
	t.Run("exists", func(t *testing.T) {
		t.Parallel()
		assert.True(t, util.PathExists(tf.Name()))
	})
	t.Run("notexists", func(t *testing.T) {
		t.Parallel()
		ne := filepath.Join(td, "notathing")
		assert.False(t, util.PathExists(ne))
	})
}

func Test_IsASN(t *testing.T) {
	type CaseT struct {
		bool
		string
	}
	cases := []CaseT{
		{
			true,
			"14525",
		},
		{
			false,
			"192.0.2.1",
		},
		{
			true,
			"AS14525",
		},
		{
			true,
			"as14525",
		},
		{
			false,
			"test",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.string, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.bool, util.IsASN(c.string), c.string)
		})
	}
}

func Test_IsIP(t *testing.T) {
	type CaseT struct {
		bool
		string
	}
	cases := []CaseT{
		{
			false,
			"AS14525",
		},
		{
			true,
			"192.0.2.1",
		},
		{
			true,
			"192.0.2.0/24",
		},
		{
			true,
			"192.0.2.1/24",
		},
		{
			true,
			"2001:db8::1",
		},
		{
			true,
			"2001:db8::1/64",
		},
		{
			true,
			"2001:db8::/64",
		},
		{
			false,
			"test",
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.string, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, c.bool, util.IsIP(c.string))
		})
	}
}
