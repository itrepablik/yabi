package yabi

import (
	"sync"

	"github.com/itrepablik/timaan"
)

const (
	_defaultTimeZone = "UTC"
)

// Config list of common config variables
type Config struct {
	TimeZone string
	mu       sync.Mutex
}

// CF short hand for 'Config' struct
var CF *Config

func init() {
	CF = SetConfig(&Config{
		TimeZone: _defaultTimeZone,
	})
}

// SetConfig set the yabi common config variables
func SetConfig(c *Config) *Config {
	c.mu.Lock()
	defer c.mu.Unlock()
	return &Config{
		TimeZone: c.TimeZone,
	}
}

// NewTimaanToken creates a new timaan token and returns with a new token
// the 'expireOn' int64 must use the unix time e.g time.Now().Add(time.Minute * 30).Unix()
func NewTimaanToken(payLoad timaan.TP, expireOn int64) (string, error) {
	rt := timaan.RandomToken()
	tok := timaan.TK{
		TokenKey: rt,
		Payload:  payLoad,
		ExpireOn: expireOn,
	}
	newToken, err := timaan.GenerateToken(rt, tok)
	if err != nil {
		return "", err
	}
	return newToken, nil
}
