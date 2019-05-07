package gateway_test

import (
	"testing"

	"github.com/danielbintar/angel/server/gateway"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	assert.NotNil(t, gateway.NewServer())
}
