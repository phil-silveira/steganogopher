package common

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestText2BitArrayConversion(t *testing.T) {
	assert := assert.New(t)
	expected := "This should be the end message"

	res := ConvertText2Bits("This should be the end message")
	result := ConvertBitArray2Text(res)

	assert.Equal(expected, result)
}

func TestAssertNoErrorShouldNoRaiseError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Should not raise error")
		}
		AssertNoError(nil)
	}()
}

func TestAssertNoErrorShouldRaiseError(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Should raise error")
		}
	}()
	AssertNoError(errors.New("random error"))
}

func TestAssertNotEmptyShouldNoRaiseError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Should not raise error")
		}
		AssertNotEmpty("string 1", "")
	}()
}

func TestAssertNotEmptyShouldRaiseError(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Should raise error")
		}
	}()
	AssertNotEmpty("", "field 'name' is required")
}
