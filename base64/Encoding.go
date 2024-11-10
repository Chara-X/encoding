package base64

import "encoding/base64"

type Encoding struct {
	enc     *base64.Encoding
	encoder string
	decoder map[byte]byte
}

func NewEncoding(encoder string) *Encoding {
	if Reference {
		return &Encoding{enc: base64.NewEncoding(encoder)}
	}
	var decoder = make(map[byte]byte)
	for i, v := range encoder {
		decoder[byte(v)] = byte(i)
	}
	return &Encoding{encoder: encoder, decoder: decoder}
}
func (enc *Encoding) EncodedLen(n int) int {
	if Reference {
		return enc.enc.EncodedLen(n)
	}
	if n%3 != 0 {
		panic("n isn't multiple of 3")
	}
	return n / 3 * 4
}
func (enc *Encoding) Encode(dst, src []byte) {
	if Reference {
		enc.enc.Encode(dst, src)
		return
	}
	for i, j := 0, 0; i < len(src); i, j = i+3, j+4 {
		var v = uint(src[i+0])<<16 | uint(src[i+1])<<8 | uint(src[i+2])
		dst[j+0] = enc.encoder[v>>18&0x3F]
		dst[j+1] = enc.encoder[v>>12&0x3F]
		dst[j+2] = enc.encoder[v>>6&0x3F]
		dst[j+3] = enc.encoder[v&0x3F]
	}
}
func (enc *Encoding) DecodedLen(n int) int {
	if Reference {
		return enc.enc.DecodedLen(n)
	}
	if n%4 != 0 {
		panic("n isn't multiple of 4")
	}
	return n / 4 * 3
}
func (enc *Encoding) Decode(dst, src []byte) (n int, err error) {
	if Reference {
		return enc.enc.Decode(dst, src)
	}
	var i, j = 0, 0
	for ; i < len(src); i, j = i+4, j+3 {
		var v = uint32(enc.decoder[src[i+0]])<<18 | uint32(enc.decoder[src[i+1]])<<12 | uint32(enc.decoder[src[i+2]])<<6 | uint32(enc.decoder[src[i+3]])
		dst[j+0] = byte(v >> 16 & 0xFF)
		dst[j+1] = byte(v >> 8 & 0xFF)
		dst[j+2] = byte(v & 0xFF)
	}
	return j, nil
}
