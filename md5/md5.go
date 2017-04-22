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
var Table = [65]uint32{
	0, 3614090360, 3905402710, 606105819, 3250441966, 4118548399, 1200080426,
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

//we have pregenerated table on top
func GenerateTable() (table [64]uint32) {
	for i := 0; i < 64; i++ {
		table[i] = uint32((1 << 32) * math.Abs(math.Sin(float64(i+1))))
	}

	return
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
	binary.LittleEndian.PutUint32(inputLen, uint32(8*len(input)))

	padded = append(padded, inputLen...)
	padded = append(padded, 0x0, 0x0, 0x0, 0x0)

	return padded
}

//auxiliary fns
func ff(x, y, z uint32) uint32 {
	return (x & y) | (^x & z)
}

func fg(x, y, z uint32) uint32 {
	return (x & z) | (y & ^z)
}

func fh(x, y, z uint32) uint32 {
	return x ^ y ^ z
}

func fi(x, y, z uint32) uint32 {
	return y | (x | ^z)
}

// not SSD one, bit rotate for short.
func bitrot32(a uint32, s uint) uint32 {
	return a<<s | a>>(32-s)
}

//Step 4
func Digest(input []byte) (res [4]uint32) {
	uintInput := Bytes2Uints(input)

	a := A
	b := B
	c := C
	d := D

	for i := 0; i < len(uintInput)/16; i++ {
		var X [16]uint32

		for j := 0; j < 16; j++ {
			X[j] = uintInput[i*16+j]
		}

		aa := a
		bb := b
		cc := c
		dd := d

		r1 := func(a, b, c, d uint32, k, s, i uint) uint32 {
			return b + bitrot32(a+ff(b, c, d)+X[k]+Table[i], s)
		}

		r2 := func(a, b, c, d uint32, k, s, i uint) uint32 {
			return b + bitrot32(a+fg(b, c, d)+X[k]+Table[i], s)
		}

		r3 := func(a, b, c, d uint32, k, s, i uint) uint32 {
			return b + bitrot32(a+fh(b, c, d)+X[k]+Table[i], s)
		}

		r4 := func(a, b, c, d uint32, k, s, i uint) uint32 {
			return b + bitrot32(a+fi(b, c, d)+X[k]+Table[i], s)
		}

		// Round 1
		a = r1(a, b, c, d, 0, 7, 1)
		d = r1(d, a, b, c, 1, 12, 2)
		c = r1(c, d, a, b, 2, 17, 3)
		b = r1(b, c, d, a, 3, 22, 4)

		a = r1(a, b, c, d, 4, 7, 5)
		d = r1(d, a, b, c, 5, 12, 6)
		c = r1(c, d, a, b, 6, 17, 7)
		b = r1(b, c, d, a, 7, 22, 8)

		a = r1(a, b, c, d, 8, 7, 9)
		d = r1(d, a, b, c, 9, 12, 10)
		c = r1(c, d, a, b, 10, 17, 11)
		b = r1(b, c, d, a, 11, 22, 12)

		a = r1(a, b, c, d, 12, 7, 13)
		d = r1(d, a, b, c, 13, 12, 14)
		c = r1(c, d, a, b, 14, 17, 15)
		b = r1(b, c, d, a, 15, 22, 16)

		// Round 2
		a = r2(a, b, c, d, 1, 5, 17)
		d = r2(d, a, b, c, 6, 9, 18)
		c = r2(c, d, a, b, 11, 14, 19)
		b = r2(b, c, d, a, 0, 20, 20)

		a = r2(a, b, c, d, 5, 5, 21)
		d = r2(d, a, b, c, 10, 9, 22)
		c = r2(c, d, a, b, 15, 1, 23)
		b = r2(b, c, d, a, 4, 20, 24)

		a = r2(a, b, c, d, 9, 5, 25)
		d = r2(d, a, b, c, 14, 9, 26)
		c = r2(c, d, a, b, 3, 14, 27)
		b = r2(b, c, d, a, 8, 20, 28)

		a = r2(a, b, c, d, 13, 5, 29)
		d = r2(d, a, b, c, 2, 9, 30)
		c = r2(c, d, a, b, 7, 14, 31)
		b = r2(b, c, d, a, 12, 20, 32)

		// Round 3
		a = r3(a, b, c, d, 5, 4, 33)
		d = r3(d, a, b, c, 8, 11, 34)
		c = r3(c, d, a, b, 11, 16, 35)
		b = r3(b, c, d, a, 14, 23, 36)

		a = r3(a, b, c, d, 1, 4, 37)
		d = r3(d, a, b, c, 4, 11, 38)
		c = r3(c, d, a, b, 7, 16, 39)
		b = r3(b, c, d, a, 10, 23, 40)

		a = r3(a, b, c, d, 13, 4, 41)
		d = r3(d, a, b, c, 0, 11, 42)
		c = r3(c, d, a, b, 3, 16, 43)
		b = r3(b, c, d, a, 6, 23, 44)

		a = r3(a, b, c, d, 9, 4, 45)
		d = r3(d, a, b, c, 12, 11, 46)
		c = r3(c, d, a, b, 15, 16, 47)
		b = r3(b, c, d, a, 2, 23, 48)

		// Round 4
		a = r4(a, b, c, d, 0, 6, 49)
		d = r4(d, a, b, c, 7, 10, 50)
		c = r4(c, d, a, b, 14, 15, 51)
		b = r4(b, c, d, a, 5, 21, 52)

		a = r4(a, b, c, d, 12, 6, 53)
		d = r4(d, a, b, c, 3, 10, 54)
		c = r4(c, d, a, b, 10, 15, 55)
		b = r4(b, c, d, a, 1, 21, 56)

		a = r4(a, b, c, d, 8, 6, 57)
		d = r4(d, a, b, c, 15, 10, 58)
		c = r4(c, d, a, b, 6, 15, 59)
		b = r4(b, c, d, a, 13, 21, 60)

		a = r4(a, b, c, d, 4, 6, 61)
		d = r4(d, a, b, c, 11, 10, 62)
		c = r4(c, d, a, b, 2, 15, 63)
		b = r4(b, c, d, a, 9, 21, 64)

		a = a + aa
		b = b + bb
		c = c + cc
		d = d + dd

	}

	res[0] = a
	res[1] = b
	res[2] = c
	res[3] = d

	return
}

// Step Hex
func Uint32hex(input [4]uint32) (bytes [16]byte) {
	binary.LittleEndian.PutUint32(bytes[0:4], input[0])
	binary.LittleEndian.PutUint32(bytes[4:8], input[1])
	binary.LittleEndian.PutUint32(bytes[8:12], input[2])
	binary.LittleEndian.PutUint32(bytes[12:16], input[3])

	return bytes
}

func Bytes2Uints(input []byte) []uint32 {
	uints := make([]uint32, len(input)/4)

	for i := 0; i < len(uints); i++ {
		uints[i] = binary.LittleEndian.Uint32(input[i*4 : (i+1)*4])
	}

	return uints
}
