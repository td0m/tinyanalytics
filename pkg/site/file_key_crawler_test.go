package site

import (
	"testing"
)

func TestFileKeyCrawler_GetDomainSignature(t *testing.T) {
	d := NewFileKeyCrawler()

	t.Run("works for a registered domain", func(t *testing.T) {
		key, err := d.GetKey("tdom.dev")
		if err != nil {
			t.Error()
		}
		if key != "works" {
			t.Errorf("should be %s but given %s", "works", key)
		}
	})
	t.Run("throws an error when unregistered", func(t *testing.T) {
		_, err := d.GetKey("tdomtdomtdom.dev")
		if err == ErrKeyFileNotFound {
			t.Error()
		}
	})
}
