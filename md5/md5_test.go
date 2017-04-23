package md5_test

import "encoding/hex"
import "github.com/nodar-chkuaselidze/go-leat/md5"
import "testing"

func TestRFCSuite(t *testing.T) {
	tests := []struct {
		src []byte
		hex string
	}{
		{[]byte(""), "d41d8cd98f00b204e9800998ecf8427e"},
		{[]byte("a"), "0cc175b9c0f1b6a831c399e269772661"},
		{[]byte("abc"), "900150983cd24fb0d6963f7d28e17f72"},
		{[]byte("message digest"), "f96b697d7cb7938d525a2f31aaf161d0"},
		{[]byte("abcdefghijklmnopqrstuvwxyz"), "c3fcd3d76192e4007dfb496cca67e13b"},
		{[]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"), "d174ab98d277d9f5a5611c2c9f419d9f"},
		{[]byte("12345678901234567890123456789012345678901234567890123456789012345678901234567890"), "57edf4a22be3c955ac49da2e2107b67a"},
	}

	for _, test := range tests {
		sum := md5.Sum(test.src)
		hexstr := hex.EncodeToString(sum[:])
		if test.hex != hexstr {
			t.Errorf("md5(%s) should be %s, is %s", test.src, test.hex, hexstr)
		}
	}
}
