package data

import (
	"fmt"
	"strconv"
)

// Runtime has a MarshalJSON() method and is used to return the value in (%d mins) format
type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
