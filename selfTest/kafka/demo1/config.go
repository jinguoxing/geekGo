package main

import (
	"strings"
	"reflect"
	"fmt"

)


type ConfigValue interface{}
type ConfigMap map[string]ConfigValue



func (m ConfigMap) clone() ConfigMap {
	m2 := make(ConfigMap)
	for k, v := range m {
		m2[k] = v
	}
	return m2
}

func (m ConfigMap) SetKey(key string, value ConfigValue) error {
	if strings.HasPrefix(key, "{topic}.") {
		_, found := m["default.topic.config"]
		if !found {
			m["default.topic.config"] = ConfigMap{}
		}
		m["default.topic.config"].(ConfigMap)[strings.TrimPrefix(key, "{topic}.")] = value
	} else {
		m[key] = value
	}

	return nil
}

// Set implements flag.Set (command line argument parser) as a convenience
// for `-X key=value` config.
func (m ConfigMap) Set(kv string) error {
	i := strings.Index(kv, "=")
	if i == -1 {
		return fmt.Errorf("Expected key=value")
	}

	k := kv[:i]
	v := kv[i+1:]

	return m.SetKey(k, v)
}


func (m ConfigMap) get(key string, defval ConfigValue) (ConfigValue, error) {
	if strings.HasPrefix(key, "{topic}.") {
		defconfCv, found := m["default.topic.config"]
		if !found {
			return defval, nil
		}
		return defconfCv.(ConfigMap).get(strings.TrimPrefix(key, "{topic}."), defval)
	}

	v, ok := m[key]
	if !ok {
		return defval, nil
	}

	if defval != nil && reflect.TypeOf(defval) != reflect.TypeOf(v) {
		//return nil, newErrorFromString(ErrInvalidArg, fmt.Sprintf("%s expects type %T, not %T", key, defval, v))
		return nil, fmt.Errorf("%s expects type %T, not %T", key, defval, v)
	}


	return v, nil
}

func (m ConfigMap) Get(key string, defval ConfigValue) (ConfigValue, error) {
	return m.get(key, defval)
}

