package lconnect

// @author  Mikhail Kirillov <mikkirillov@yandex.ru>
// @version 1.001
// @date    2019-10-16

import (
	"time"
)

var (
	KEEP_ALIVE              = time.Second * 30
	RECONNECT_AFTER_SECONDS = int64(3)
)
