package tools

import (
	"proto_buffer_example/server/third-party/antnet"
	"reflect"
	"strconv"
)

func GetInterfaceType(data interface{}) string {
	if t := reflect.TypeOf(data); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}
func Interface2String(value interface{}) (string, bool) {
	success := false
	switch value.(type) {
	case string:
		success = true
		return value.(string), success
	}
	return "", false
}

func Interface2uint64(value interface{}) (uint64, bool) {
	success := false
	switch value.(type) {
	case uint64:
		success = true
		return value.(uint64), success
	case string:
		i, err := strconv.ParseUint(value.(string), 10, 64)
		if err == nil {
			success = true
		}
		return i, success
	}
	return 0, false
}

func Interface2int64(value interface{}) (int64, bool) {
	success := false
	switch value.(type) {
	case int64:
		success = true
		return value.(int64), success
	case string:
		i := antnet.Atoi64(value.(string))
		if i != 0 {
			success = true
		}
		return i, success
	}
	return 0, false
}

func Interface2int32(value interface{}) (int32, bool) {
	success := false
	switch value.(type) {
	case int32:
		success = true
		return value.(int32), success
	case string:
		i := antnet.Atoi(value.(string))
		if i != 0 {
			success = true
		}
		return int32(i), success
	}
	return 0, false
}

func Interface2int(value interface{}) (int, bool) {
	success := false
	switch value.(type) {
	case int:
		success = true
		return value.(int), success
	case string:
		i := antnet.Atoi(value.(string))
		if i != 0 {
			success = true
		}
		return i, success
	}
	return 0, false
}

func Interface2float64(value interface{}) (float64, bool) {
	success := false
	switch value.(type) {
	case float64:
		success = true
		return value.(float64), success
	case string:
		i := antnet.Atof64(value.(string))
		if i != 0 {
			success = true
		}
		return i, success
	}
	return 0, false
}

func Interface2Bool(value interface{}) (bool, bool) {
	success := false
	switch value.(type) {
	case bool:
		success = true
		return value.(bool), success
	}
	return false, false
}
