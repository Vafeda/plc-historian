package models

import (
	"time"

	"github.com/gopcua/opcua/ua"
)

type Sample struct {
	ItemId    uint32
	ItemTS    time.Time
	RegTS     time.Time
	ValInt    *int64
	ValDouble *float64
	ValBool   *bool
	ValStr    *string
	Quality   ua.StatusCode
}
