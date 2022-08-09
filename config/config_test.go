package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHTTPConfig(t *testing.T) {
	assert := assert.New(t)
	hc := &HTTPConfig{}
	rht, rt, wt := hc.Timeouts()
	assert.Equal(rht, 10*time.Second)
	assert.Equal(rt, time.Minute)
	assert.Equal(wt, time.Minute)

	hc.ReadHeaderTimeoutMS = 10
	hc.ReadTimeoutMS = 1000
	hc.WriteTimeoutMS = 1000

	rht, rt, wt = hc.Timeouts()
	assert.Equal(rht, 10*time.Millisecond)
	assert.Equal(rt, time.Second)
	assert.Equal(wt, time.Second)

	assert.Equal("127.0.0.1:8080", hc.ListenAddr())
	hc.Listen = ":9090"
	assert.Equal("127.0.0.1:9090", hc.ListenAddr())

	hc.Listen = "127.0.0.2"
	assert.Equal("127.0.0.2:8080", hc.ListenAddr())

	hc.Listen = "127.0.0.1:4000"

	os.Setenv("ANGLE_HTTP_HOST", "127.0.0.3")
	assert.Equal("127.0.0.3:4000", hc.ListenAddr())

	os.Setenv("ANGLE_HTTP_PORT", "3000")
	assert.Equal("127.0.0.3:3000", hc.ListenAddr())
}
