package base64

import "encoding/base64"

var Reference = false
var (
	StdEncoding = &Encoding{enc: base64.StdEncoding}
	URLEncoding = &Encoding{enc: base64.URLEncoding}
)
