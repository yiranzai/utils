package structutil

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// 复制一个Model
func CopyModel(from, to interface{}) error {
	if from == nil || to == nil {
		return nil
	}

	var out interface{}

	// convert all from object to []map[string]interface{}
	// bool for JSON booleans,
	// float64 for JSON numbers,
	// string for JSON strings, and
	// nil for JSON null.
	data, err := json.Marshal(from)
	if err != nil {
		return fmt.Errorf("copy model marshal failed: %w", err)
	}
	if err := json.Unmarshal(data, &out); err != nil {
		return fmt.Errorf("copy model unmarshal to map failed: %w", err)
	}

	// transform time field
	transform(out)

	// convert to json data
	data, err = json.Marshal(out)
	if err != nil {
		return fmt.Errorf("copy model marsham map to data failed: %w", err)
	}

	// convert to struct data
	if err := json.Unmarshal(data, to); err != nil {
		return fmt.Errorf("copy model to target failed: %w", err)
	}
	return nil
}

func transformTimeToUnix(k string, v interface{}) interface{} {
	k = strings.ToLower(k)
	if !strings.HasSuffix(k, "at") && !strings.HasSuffix(k, "time") {
		return v
	}
	str, ok := v.(string)
	if !ok {
		return v
	}
	if len(str) < (len(time.RFC3339)-5) || len(time.RFC3339Nano) < len(str) {
		return v
	}
	t, err := time.ParseInLocation(time.RFC3339Nano, str, time.Local)
	if err != nil {
		return v
	}
	if t.IsZero() {
		return 0
	}
	return t.Unix()
}

func transformUnixToTime(k string, v interface{}) (interface{}, bool) {
	k = strings.ToLower(k)
	if !strings.HasSuffix(k, "at") && !strings.HasSuffix(k, "time") {
		return v, false
	}
	value, ok := v.(float64)
	if !ok {
		return v, false
	}
	return time.Unix(int64(value), 0).Format(time.RFC3339Nano), true
}

func transform(v interface{}) {
	if v, ok := v.([]interface{}); ok {
		for _, v := range v {
			transform(v)
		}
	}
	if m, ok := v.(map[string]interface{}); ok {
		transformMap(m)
	}
}

func transformMap(m map[string]interface{}) {
	for k, v := range m {
		// decode sub objects
		if mm, ok := v.(map[string]interface{}); ok {
			transformMap(mm)
		}
		// decode array objects
		if v, ok := v.([]interface{}); ok {
			transform(v)
		}
		// decode single field
		var ok bool
		m[k], ok = transformUnixToTime(k, m[k])
		if !ok {
			m[k] = transformTimeToUnix(k, m[k])
		}
	}
}
