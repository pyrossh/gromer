package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCss(t *testing.T) {
	assert.Equal(t, "123", Css("123").classes)
}
