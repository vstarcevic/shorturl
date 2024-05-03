package app

import (
	"testing"
)

func TestBase58Encode(t *testing.T) {

	// Arange
	var num = []int64{10, 15, 58, 100, 10001, 1000000}
	should := []string{"B", "G", "21", "2j", "3yS", "68GP"}

	for i := range num {

		// Act
		result := EncodeBase58(num[i])

		// Assert
		if result != should[i] {
			t.Fatalf(`Result not as expected. Got "%v", and it should be %v`, result, should[i])
		}
	}
}
