package md5

import (
	"encoding/binary"
	"math"
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

//Don't compute table all the time
var Table = [64]uint32{
	3614090360, 3905402710, 606105819, 3250441966, 4118548399, 1200080426,
	2821735955, 4249261313, 1770035416, 2336552879, 4294925233, 2304563134,
	1804603682, 4254626195, 2792965006, 1236535329, 4129170786, 3225465664,
	643717713, 3921069994, 3593408605, 38016083, 3634488961, 3889429448,
	568446438, 3275163606, 4107603335, 1163531501, 2850285829, 4243563512,
	1735328473, 2368359562, 4294588738, 2272392833, 1839030562, 4259657740,
	2763975236, 1272893353, 4139469664, 3200236656, 681279174, 3936430074,
	3572445317, 76029189, 3654602809, 3873151461, 530742520, 3299628645,
	4096336452, 1126891415, 2878612391, 4237533241, 1700485571, 2399980690,
	4293915773, 2240044497, 1873313359, 4264355552, 2734768916, 1309151649,
	4149444226, 3174756917, 718787259, 3951481745,
}

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

//auxiliary fns
func F(x, y, z uint32) uint32 {
	return (x & y) | (^x & z)
}

func G(x, y, z uint32) uint32 {
	return (x & z) | (y & ^z)
}

func H(x, y, z uint32) uint32 {
	return x ^ y ^ z
}

func I(x, y, z uint32) uint32 {
	return y | (x | ^z)
}

//we have pregenerated table on top
func GenerateTable() (table [64]uint32) {
	for i := 0; i < 64; i++ {
		table[i] = uint32((1 << 32) * math.Abs(math.Sin(float64(i+1))))
	}

	return
}

//Step 4
