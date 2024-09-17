package strings

import (
	"strings"
	"testing"
)

const indexMessage = "%s (parte %s) - índices: esperado (%d) <> encontrado (%d)."

type TestType struct {
	text string
	part string
	expected int
}

func TestIndex(t *testing.T) {
	tests := []TestType {
		{"Enzo's test", "Enzo", 0},
		{"", "", 0},
		{"Ops", "ops", -1},
		{"Enzo", "z", 2},
	}

	for _, test := range tests {
		t.Logf("Test: %v", test)
		current := strings.Index(test.text, test.part)

		if current != test.expected {
			t.Errorf(indexMessage, test.text, test.part, test.expected, current)
		}
	}
}
