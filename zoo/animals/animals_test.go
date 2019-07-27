package animals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestElephantFeed(t *testing.T) {
	expected := "Grass"
	actual := ElephantFeed()
	assert.Equal(t, expected, actual)
}

func TestMonkeyFeed(t *testing.T) {
	expected := "Banana"
	actual := MonkeyFeed()
	assert.Equal(t, expected, actual)
}

func TestRabbitFeed(t *testing.T) {
	expected := "Carrot"
	actual := RabbitFeed()
	assert.Equal(t, expected, actual)
}
