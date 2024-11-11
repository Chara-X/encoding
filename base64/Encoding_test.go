package base64

import "fmt"

func ExampleEncoding() {
	var plaintext = []byte("Hello, World")
	var enc = NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	var encoded = make([]byte, enc.EncodedLen(len(plaintext)))
	enc.Encode(encoded, plaintext)
	plaintext = make([]byte, enc.DecodedLen(len(encoded)))
	var n, _ = enc.Decode(plaintext, encoded)
	fmt.Println(n)
	fmt.Println(string(plaintext))
	// Output:
	// 12
	// Hello, World
}
