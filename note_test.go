package main

import (
    "github.com/stretchr/testify/assert"
    "testing"
)
func TestNoteTitle(t *testing.T) {
	n := &note{content: "this note"}
	assert.Equal(t, "this note", n.title)

    n = &note{content: "line 1\nline 2\n"}
    assert.Equal(t, "line 1", n.title)
}
