package tests_test

import (
	"testing"

	"github.com/dorenproctor/easyexec"
	"github.com/stretchr/testify/assert"
)

func TestRunEcho(t *testing.T) {
	got := easyexec.Run("echo", "foo")
	assert.NoError(t, got.Err)
	assert.Equal(t, "foo\n", got.Stdout)
	assert.Equal(t, "", got.Stderr)
}
