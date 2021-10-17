package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search_url(t *testing.T) {
	got := searchURL("a", "b")
	want := fmt.Sprintf("%s?q=a&key=b&limit=10", tenorAPIURLbase)

	assert.Equal(t, got, want)
}
