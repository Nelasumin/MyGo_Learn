package demo01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJudge(t *testing.T) {

	isPass := Judge(70)
	assert.Equal(t, true, isPass)
}

func TestJudge2(t *testing.T) {
	isPass := Judge(50)
	assert.Equal(t, false, isPass)
}
