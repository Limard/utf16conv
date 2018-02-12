package utf16conv

import (
	"unicode/utf16"
	"unsafe"
)

func U8ToU16Array(s string) []uint16 {
	return utf16.Encode([]rune(s + "\x00"))
}

func U8ToU16Byte(u8 string) []byte {
	u16 := U8ToU16Array(u8)
	return U16ArrayToByte(u16)
}

func U16ArrayToU8(s []uint16) string {
	for i, v := range s {
		if v == 0 {
			s = s[0:i]
			break
		}
	}
	return string(utf16.Decode(s))
}

func U16ByteToU8(s []byte) string {
	u16 := ByteToU16Array(s)
	return U16ArrayToU8(u16)
}

func U16ArrayToByte(s []uint16) (o []byte) {
	o = make([]byte, 0, len(s)*2)
	for i := 0; i < len(s); i++ {
		o = append(o, byte(s[i]&0xFF), byte(s[i]>>8))
	}
	return
}

func ByteToU16Array(s []byte) (o []uint16) {
	o = make([]uint16, 0, len(s)/2)
	for i := 0; i < len(s); i += 2 {
		o = append(o, *(*uint16)(unsafe.Pointer(&s[i])))
	}
	return
}
