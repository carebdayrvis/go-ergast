package ergast

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseErgastDuration(t *testing.T) {

	e, err := time.ParseDuration("1m40s650ms")
	assert.Nil(t, err)

	a, err := parseErgastDuration("1:40.650")
	assert.Nil(t, err)

	assert.Equal(t, e, a)
}
