package ch5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// drawLine draws a horizontal line on the screen. The screen is represented
// by a byte array with each bit representing one pixel. The width is divisible
// by 8 so no byte is split across lines.
func drawLine(screen []byte, width, x1, x2, y int) []byte {
	bytesPerRow := width / 8
	byteAtRowStart := y * bytesPerRow
	minByteInRow := x1 / 8
	maxByteInRow := x2 / 8
	minByte := byteAtRowStart + minByteInRow
	maxByte := byteAtRowStart + maxByteInRow
	minByteMinBit := x1 % 8
	maxByteMaxBit := x2 % 8
	// TODO: Error handling.
	for i := minByte; i <= maxByte; i++ {
		minBit := 0
		if i == minByte {
			minBit = minByteMinBit
		}
		maxBit := 7
		if i == maxByte {
			maxBit = maxByteMaxBit
		}
		for j := minBit; j <= maxBit; j++ {
			// setBit indexes from right to left so we need to reverse the bits
			// before and after.
			screen[i] = reverseBits(byte(setBit(int(reverseBits(screen[i])), j, 1)))
		}
	}
	return screen
}

// reverseBits reverses the bits from right-to-left to left-to-right.
func reverseBits(x byte) byte {
	result := byte(0)
	for i := 0; i < 8; i++ {
		bit := x & 1
		x = x >> 1
		result = result << 1
		result = result | bit
	}
	return result
}

func TestDrawLine1(t *testing.T) {
	screen := []byte{
		0b00000000, 0b00000000,
		0b00000000, 0b00000000,
		0b00000000, 0b00000000,
		0b00000000, 0b00000000,
	}
	expected := []byte{
		0b00000000, 0b00000000,
		0b00000111, 0b11100000,
		0b00000000, 0b00000000,
		0b00000000, 0b00000000,
	}
	assert.Equal(t, expected, drawLine(screen, 16, 5, 10, 1))
}

func TestDrawLine2(t *testing.T) {
	screen := []byte{
		0b00000000, 0b00000000,
		0b00000000, 0b00000000,
		0b00000000, 0b00000000,
		0b00000000, 0b00000000,
	}
	expected := []byte{
		0b00000000, 0b00000000,
		0b00000110, 0b00000000,
		0b00000000, 0b00000000,
		0b00000000, 0b00000000,
	}
	assert.Equal(t, expected, drawLine(screen, 16, 5, 6, 1))
}

func TestDrawLine3(t *testing.T) {
	screen := []byte{
		0b00000000, 0b00000000,
		0b00000000, 0b00000000,
		0b00000000, 0b00000000,
		0b00000000, 0b00000000,
	}
	expected := []byte{
		0b00000000, 0b00000000,
		0b11111111, 0b11111111,
		0b00000000, 0b00000000,
		0b00000000, 0b00000000,
	}
	assert.Equal(t, expected, drawLine(screen, 16, 0, 15, 1))
}

func TestReverseBits(t *testing.T) {
	assert.Equal(t, byte(0b01011000), reverseBits(0b00011010))
}
