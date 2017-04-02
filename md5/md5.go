package md5

import (
	"encoding/binary"
)

// MD5 Implementation
// Goal isn't to have fully functioning implementation, nor speed
// And MUST NOT be used in production or anywhere near that

// Limits:
//  Input can't be more than 2^32, this limitation is imposed
// by golang []bytes

const (
	A = uint32(0x67452301)
	B = uint32(0xefcdab89)
	C = uint32(0x98badcfe)
	D = uint32(0x10325476)
)

//Step 1
func AddPadding(input []byte) []byte {
	padded := make([]byte, len(input))
	copy(padded, input)
	padded = append(padded)

	left := len(padded) % 64

	if left > 56 {
		left = 56 + (64 - left)
	} else {
		left = 56 - left
	}

	add := make([]byte, left)
	add[0] = add[0] | (1 << 7)
	padded = append(padded, add...)

	return padded
}

//Step 2
func AppendLength(input, padded []byte) []byte {
	inputLen := make([]byte, 4)
	binary.LittleEndian.PutUint32(inputLen, uint32(len(input)))

	padded = append(padded, inputLen...)
	padded = append(padded, 0x0, 0x0, 0x0, 0x0)

	return padded
}
