package url_config_test

import (
	"testing"

	"github.com/danielbintar/angel/server/gateway/url_config"

	"github.com/stretchr/testify/assert"
)

func TestInstance(t *testing.T) {
	assert.NotNil(t, url_config.Instance("./"))
	assert.NotNil(t, url_config.Instance("../url_factory/valid/"))
	assert.NotNil(t, url_config.Instance("../url_factory/no_prefix/"))
	assert.Panics(t, func() { url_config.Instance("../url_factory/imagination/") })
	assert.Panics(t, func() { url_config.Instance("../url_factory/invalid_format/") })
}
