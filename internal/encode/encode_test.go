package encode

import "testing"

func TestEncodeTextIntoPNGShouldSucceed(t *testing.T) {
	Encode("hidden message XD", "../../examples/le-petit-prince.png", "../../examples/le-petit-prince.test.png")
}

func TestEncodeTextWithoutPassOutputPathIntoPNGShouldSucceed(t *testing.T) {
	Encode("hidden message XD", "../../examples/le-petit-prince.png", "")
}
