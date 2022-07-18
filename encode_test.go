package steganogopher

import "testing"

func TestEncodeTextIntoPNGShouldSucceed(t *testing.T) {
	Encode("hidden message XD", "./examples/avatar.png", "./examples/avatar_test.png")
}

func TestEncodeTextWithoutPassOutputPathIntoPNGShouldSucceed(t *testing.T) {
	Encode("hidden message XD", "./examples/avatar.png", "")
}
