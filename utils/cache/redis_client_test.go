package cache

import (
	"testing"
)

func TestNewRedisCache(t *testing.T) {
	a := NewRedisCache()
	a.LPush("name", "appple")
	a.BRpop("name")
}
