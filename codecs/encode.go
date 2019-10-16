package codecs

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-10-16

import (
	"github.com/belfinor/pack"
)

type Encode struct {
}

func (e *Encode) Write(flags int32, data []byte) []byte {

	return pack.Encode(int32(len(data)+4), flags, data)
}
