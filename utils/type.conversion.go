package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// ToInterfaceArray converts interface{} to []interface{} and never returns nil
func ToInterfaceArray(i1 interface{}) []interface{} {
	if i1 == nil {
		return []interface{}{}
	}
	switch i2 := i1.(type) {
	default:
		return []interface{}{}
	case []interface{}:
		return i2
	}
	return []interface{}{}
}

// ToInterfaceMap converts interface{} to map[string]interface{} and never returns nil
func ToInterfaceMap(i1 interface{}) map[string]interface{} {
	if i1 == nil {
		return map[string]interface{}{}
	}
	switch i2 := i1.(type) {
	case map[string]interface{}:
		return i2
	default:
		return map[string]interface{}{}
	}
	return map[string]interface{}{}
}

// ToStringArray converts interface{} to []string and never returns nil
func ToStringArray(i1 interface{}) []string {
	if i1 == nil {
		return []string{}
	}
	switch i2 := i1.(type) {
	default:
		return []string{fmt.Sprint(i2)}
	case []string:
		return i2
	case []interface{}:
		var ss []string
		for _, i3 := range i2 {
			ss = append(ss, ToString(i3))
		}
		return ss
	}
	return []string{}
}

// ToStringMap converts interface{} to map[string]string and never returns nil
func ToStringMap(i1 interface{}) map[string]string {
	switch i2 := i1.(type) {
	case map[string]interface{}:
		m1 := map[string]string{}
		for k, v := range i2 {
			m1[k] = ToString(v)
		}
		return m1
	case map[string]string:
		return i2
	default:
		return map[string]string{}
	}
}

// ToString converts interface{} to string
func ToString(i1 interface{}) string {
	if i1 == nil {
		return ""
	}

	switch i2 := i1.(type) {
	default:
		return fmt.Sprint(i2)
	case bool:
		if i2 {
			return "true"
		} else {
			return "false"
		}
	case string:
		return i2
	case *bool:
		if i2 == nil {
			return ""
		}

		if *i2 {
			return "true"
		} else {
			return "false"
		}
	case *string:
		if i2 == nil {
			return ""
		}
		return *i2
	case *json.Number:
		return i2.String()
	case json.Number:
		return i2.String()
	}

	return ""
}

// ToInt64 converts interface{} to int64
func ToInt64(i1 interface{}) int64 {
	if i1 == nil {
		return 0
	}

	switch i2 := i1.(type) {
	default:
		i3, _ := strconv.ParseInt(ToString(i2), 10, 64)
		return i3
	case *json.Number:
		i3, _ := i2.Int64()
		return i3
	case json.Number:
		i3, _ := i2.Int64()
		return i3
	case int64:
		return i2
	case float64:
		return int64(i2)
	case float32:
		return int64(i2)
	case uint64:
		return int64(i2)
	case int:
		return int64(i2)
	case uint:
		return int64(i2)
	case bool:
		if i2 {
			return 1
		} else {
			return 0
		}
	case *bool:
		if i2 == nil {
			return 0
		}

		if *i2 {
			return 1
		} else {
			return 0
		}
	}

	return 0
}

// ToBool converts interface{} to bool
func ToBool(i1 interface{}) bool {
	if i1 == nil {
		return false
	}

	switch i2 := i1.(type) {
	default:
		return false
	case bool:
		return i2
	case string:
		return i2 == "true"
	case int:
		return i2 != 0
	case *bool:
		if i2 == nil {
			return false
		}
		return *i2
	case *string:
		if i2 == nil {
			return false
		}
		return *i2 == "true"
	case *int:
		if i2 == nil {
			return false
		}
		return *i2 != 0
	}
	return false
}
