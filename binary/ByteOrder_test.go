package binary

import "fmt"

func ExampleBigEndian() {
	var buf = make([]byte, 12)
	BigEndian.PutUint32(buf, 1)
	BigEndian.PutUint64(buf[4:], 2)
	fmt.Println(BigEndian.Uint32(buf))
	fmt.Println(BigEndian.Uint64(buf[4:]))
	// Output:
	// 1
	// 2
}
