package daemon

import (
	"bytes"
	"testing"

	"github.com/stvp/assert"
)

func TestLogs(t *testing.T) {
	startForTest()
	var stream bytes.Buffer
	err := Logs(&stream)
	assert.Nil(t, err)
	assert.NotNil(t, stream)
}