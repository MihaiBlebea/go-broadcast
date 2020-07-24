package hello

import (
	"fmt"
	"testing"
)

func TestBroadcastMessage(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		template string
		expected interface{}
	}{
		{"Mihai", 30, "My name is %s and I am %d yo", "My name is Mihai and I am 30 yo"},
		{"Florin", 0, "My name is %s and I am %d yo", ErrNoAge},
	}

	for _, test := range tests {
		serv := New()

		testName := fmt.Sprintf(
			"Name_%v_age_%v_template_%v_expected_%v",
			test.name,
			test.age,
			test.template,
			test.expected,
		)

		t.Run(testName, func(t *testing.T) {
			msg, err := serv.BroadcastMessage(test.name, test.age, test.template)
			if err != nil {
				if err != test.expected {
					t.Errorf("Expected %v, got %v", test.expected, err)
				}
				return
			}

			if msg != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, msg)
			}
		})
	}
}
