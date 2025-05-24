package data

import (
	"fmt"
	"strconv"
)

// we are doing this because go looks for this method before using it default marshaling method

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	// simply converts the int32 to string
	jsonValue := fmt.Sprintf("%d mins", r)

	// this adds double quotes to the string to make it proper json string
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
