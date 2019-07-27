package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppName(t *testing.T) {
	expected := "Zoo Application"
	actual := AppName()
	assert.Equal(t, expected, actual)
}
