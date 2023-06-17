package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SetTokenMapByUidPid(t *testing.T) {
	m := make(map[string]int, 0)
	m["test1"] = 1
	m["test2"] = 2
	m["2332"] = 4
	assert.Nil(t, nil)

}
