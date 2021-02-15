package cache

import (
	"testing"
	"time"
)

func TestMap_Save(t *testing.T) {
	action := "upvote"
	action2 := "view"

	after := time.Second

	m := NewMap(after)
	new := m.Store(action)
	if !new {
		t.Fatal("expected 1st map save to be new")
	}

	new = m.Store(action)
	if new {
		t.Fatal("expected 2nd map save to not be new")
	}

	new = m.Store(action2)
	if !new {
		t.Fatal("saves to different actions should make no difference")
	}
}

func TestMap_ShouldExpire(t *testing.T) {
	action := "upvote"
	after := time.Second

	m := NewMap(after)
	m.Store(action)
	if m.Store(action) {
		t.Fatal("should return false before expiring")
	}

	time.Sleep(after)
	if !m.Store(action) {
		t.Fatal("should return true once expired")
	}
	if m.Store(action) {
		t.Fatal("second save not successful")
	}
}
