package tests_test

import (
	"fmt"
	"testing"

	"github.com/dorenproctor/easyexec"
)

func TestRunEcho(t *testing.T) {
	got := easyexec.Run("echo", "foo")
	assertNoError(t, got.Err)
	assertEqual(t, "foo\n", got.Stdout)
	assertEqual(t, "", got.Stderr)
}

func assertNoError(t *testing.T, e error, msgAndArgs ...interface{}) {
	if e != nil {
		t.Error(e, "\n", fmt.Sprint(msgAndArgs...))
	}
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}, msgAndArgs ...interface{}) {
	if expected != actual {
		t.Errorf("Not equal:\nexpected: %v\nactual  : %v\n%v",
			expected, actual, fmt.Sprint(msgAndArgs...))
	}
}
