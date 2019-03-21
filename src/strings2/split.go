package strings2

import (
	"strings"
	"errors"
	"fmt"
)

// Split 以sep为分隔标识， length为单段长度， 对origin进行分割
// 返回：分割后每一段都为length及以下的string切片
func Split(origin string, sep string, length int) ([]string, error) {

	var result = make([]string, 0)

	if len(origin) <= length {
		return append(result, origin), nil
	}

	var start, end, index int

	for len(origin[end:]) > length {
		index = strings.Index(origin[end:], sep)

		if index < 0 {
			return nil, errors.New(fmt.Sprintf("found no sep in the origin string: %s", origin[start:end]))
		}

		if index > length {
			return nil, errors.New(fmt.Sprintf("wrong length of string: %s : len=%d", origin[start:end], len(origin[start:end])))
		}

		if end-start+index >= length {
			result = append(result, origin[start:end-1])
			start = end
		}

		end = end + index + 1
	}

	result = append(result, origin[start:end-1])

	result = append(result, origin[end:])

	return result, nil
}
