package rid

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	id := New("prd", true)

	if len(id.String()) != 36 {
		t.Errorf("expected length 36, got %d (%s)", len(id.String()), id)
	}

	t.Logf("generated ID: %s", id)
}
