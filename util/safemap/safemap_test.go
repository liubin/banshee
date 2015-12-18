// Copyright 2015 Eleme Inc. All rights reserved.

package safemap

import (
	"fmt"
	"github.com/eleme/banshee/util/assert"
	"testing"
)

func TestSafeMap(t *testing.T) {
	m := New()
	// Set
	m.Set("key", "val")
	assert.Ok(t, m.Len() == 1)
	// Get
	val, ok := m.Get("key")
	assert.Ok(t, ok)
	assert.Ok(t, val == "val")
	// Items
	for k, v := range m.Items() {
		fmt.Printf("%v=>%v", k, v)
	}
	assert.Ok(t, m.Items()["key"] == "val")
	// Delete
	assert.Ok(t, m.Delete("key"))
	assert.Ok(t, !m.Delete("key1"))
}