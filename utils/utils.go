package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ToStr(value interface{}) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', 2, 32)
	case float64:
		s = strconv.FormatFloat(v, 'f', 2, 64)
	case int:
		s = strconv.FormatInt(int64(v), 10)
	case int8:
		s = strconv.FormatInt(int64(v), 10)
	case int16:
		s = strconv.FormatInt(int64(v), 10)
	case int32:
		s = strconv.FormatInt(int64(v), 10)
	case int64:
		s = strconv.FormatInt(v, 10)
	case uint:
		s = strconv.FormatUint(uint64(v), 10)
	case uint8:
		s = strconv.FormatUint(uint64(v), 10)
	case uint16:
		s = strconv.FormatUint(uint64(v), 10)
	case uint32:
		s = strconv.FormatUint(uint64(v), 10)
	case uint64:
		s = strconv.FormatUint(v, 10)
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

func JoinURLPath(path ...string) string {

	pLen := len(path)

	if pLen < 1 {
		return ""
	}

	sb := new(strings.Builder)
	for i, v := range path {
		sb.WriteString(PathCleanSep(v))
		if i < pLen-1 {
			sb.WriteString(URLSeparator)
		}
	}

	return sb.String()
}

func PathCleanSep(path string) string {
	path = strings.TrimSpace(path)

	if path == "" {
		return path
	}

	path = strings.TrimRight(path, URLSeparator)
	return strings.TrimLeft(path, URLSeparator)
}
