package address

import (
	"testing"
)

func TestNewAddress(t *testing.T) {
	addr := NewAddress(1, 2)
	if addr.Row != 1 || addr.Col != 2 {
		t.Errorf("NewAddress(1, 2) = %v; want {Row: 1, Col: 2}", addr)
	}
}
