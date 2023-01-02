package utils

import "testing"

func TestAliasID(t *testing.T) {
	id, err := AliasID("ab")

	t.Logf("id: %d, err: %v", id, err)
}
