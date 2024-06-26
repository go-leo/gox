package randx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumericPermString(t *testing.T) {
	permString := NumericString(10)
	assert.Len(t, permString, 10)
	t.Log(permString)

	permString = NumericString(15)
	assert.Len(t, permString, 15)
	t.Log(permString)

	permString = NumericString(20)
	assert.Len(t, permString, 20)
	t.Log(permString)

	permString = NumericString(21)
	assert.Len(t, permString, 21)
	t.Log(permString)

	permString = NumericString(30)
	assert.Len(t, permString, 30)
	t.Log(permString)

	permString = NumericString(39)
	assert.Len(t, permString, 39)
	t.Log(permString)

}

func TestWordString(t *testing.T) {
	permString := WordString(10)
	assert.Len(t, permString, 10)
	permString = WordString(15)
	assert.Len(t, permString, 15)
	permString = WordString(20)
	assert.Len(t, permString, 20)
	permString = WordString(21)
	assert.Len(t, permString, 21)
	permString = WordString(30)
	assert.Len(t, permString, 30)
	permString = WordString(39)
	assert.Len(t, permString, 39)
	permString = WordString(16)
	t.Log(permString)
}

func TestHexString(t *testing.T) {
	permString := HexString(10)
	assert.Len(t, permString, 10)
	permString = HexString(15)
	assert.Len(t, permString, 15)
	permString = HexString(20)
	assert.Len(t, permString, 20)
	permString = HexString(21)
	assert.Len(t, permString, 21)
	permString = HexString(30)
	assert.Len(t, permString, 30)
	permString = HexString(39)
	assert.Len(t, permString, 39)
	permString = HexString(16)
	assert.Len(t, permString, 16)
}
