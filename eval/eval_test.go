package eval

import (
	"testing"

	"github.com/evanphx/m13/builtin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vektra/neko"
)

func TestEvaluator(t *testing.T) {
	n := neko.Start(t)

	n.It("can evaluate code", func() {
		ev, err := NewEvaluator()
		require.NoError(t, err)

		val, err := ev.Eval("3")
		require.NoError(t, err)

		i, ok := val.(builtin.I64)
		require.True(t, ok)

		assert.Equal(t, builtin.I64(3), i)
	})

	n.It("runs a lambda invocation", func() {
		ev, err := NewEvaluator()
		require.NoError(t, err)

		val, err := ev.Eval(`a = x => x + 3; a(4)`)
		require.NoError(t, err)

		i, ok := val.(builtin.I64)
		require.True(t, ok)

		assert.Equal(t, builtin.I64(7), i)
	})

	n.It("runs a method call", func() {
		ev, err := NewEvaluator()
		require.NoError(t, err)

		val, err := ev.Eval(`3.add(4)`)
		require.NoError(t, err)

		i, ok := val.(builtin.I64)
		require.True(t, ok)

		assert.Equal(t, builtin.I64(7), i)
	})

	n.Meow()
}
