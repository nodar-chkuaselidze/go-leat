package leats

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/nodar-chkuaselidze/go-leat/md5"
)

func LeatMd5(length int, hash, extend, byteFormat string) (newhash string, extendWith string, err error) {
	mockbytes := make([]byte, length)
	padded := md5.AddPadding(mockbytes)
	lenPadded := md5.AppendLength(length, padded)
	padding := lenPadded[length:]

	oldmd5, err := hex.DecodeString(hash)

	if err != nil {
		return
	}

	if len(oldmd5) != 16 {
		err = errors.New("MD5 size is 16 bytes")
		return
	}

	extendBytes := []byte(extend)

	uint32s := md5.Bytes2Uints(oldmd5[:])
	A := uint32s[0]
	B := uint32s[1]
	C := uint32s[2]
	D := uint32s[3]

	cpad := md5.AddPadding(extendBytes)
	clen := md5.AppendLength(len(extend)+len(lenPadded), cpad)
	newmd5 := md5.Digest(A, B, C, D, clen)
	newmd5b := md5.Uint32Bytes(newmd5)

	newhash = hex.EncodeToString(newmd5b[:])

	extendWith = ""
	for _, b := range padding {
		extendWith += fmt.Sprintf(byteFormat, b)
	}
	extendWith = extendWith + string(extendBytes)

	return
}
