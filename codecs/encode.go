package codecs

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-10-15

import (
	"github.com/belfinor/pack"
)

type Encode struct {
}

func (e *Encode) Write(data []byte) []byte {

	return pack.Encode(int32(len(data)), data)
}
