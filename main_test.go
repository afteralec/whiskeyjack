package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeadMeta(t *testing.T) {
	hb := Head()
	hb.Meta("charset=\"UTF-8\"")
	hb.Meta("name=\"viewport\"", "content=\"width=device-width, initial-scale=1.0\"")
	result := hb.Prepare()

	require.Equal(t, result, "<head><meta charset=\"UTF-8\"/><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"/>")
}
