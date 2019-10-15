package codecs

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.000
// @date    2019-10-15

import (
	"bytes"

	"github.com/belfinor/pack"
)

type Decode struct {
	data []byte
}

func (d *Decode) Write(data []byte) [][]byte {

	var list []byte

	if d.data == nil {
		list = bytes.Join([][]byte{[]byte{}, data}, nil)
	} else {
		list = bytes.Join([][]byte{d.data, data}, nil)
	}

	var res [][]byte

	size := int32(0)

	for len(list) > 4 {
		if pack.Decode(list, &size) != nil {
			break
		}
		size = size + 4
		if len(list) > int(size) {
			res = append(res, list[4:size])
			list = list[size:]
		} else if len(list) == int(size) {
			res = append(res, list[4:])
			list = []byte{}
		} else {
			break
		}
	}

	if len(list) > 0 {
		d.data = list
	} else {
		d.data = []byte{}
	}

	return res
}
