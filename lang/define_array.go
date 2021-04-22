package lang

import (
	"fmt"

	"github.com/lmorg/murex/lang/proc/stdio"
)

// ArrayTemplate is a template function for reading arrays from marshalled data
func ArrayTemplate(marshal func(interface{}) ([]byte, error), unmarshal func([]byte, interface{}) error, read stdio.Io, callback func([]byte)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	var v interface{}
	err = unmarshal(b, &v)

	if err != nil {
		return err
	}

	switch v.(type) {
	case string:
		return readArrayByString(v.(string), callback)

	case []string:
		return readArrayBySliceString(v.([]string), callback)

	case []interface{}:
		return readArrayBySliceInterface(marshal, v.([]interface{}), callback)

	case map[string]string:
		return readArrayByMapStrStr(v.(map[string]string), callback)

	case map[string]interface{}:
		return readArrayByMapStrIface(marshal, v.(map[string]interface{}), callback)

	case map[interface{}]string:
		return readArrayByMapIfaceStr(v.(map[interface{}]string), callback)

	case map[interface{}]interface{}:
		return readArrayByMapIfaceIface(marshal, v.(map[interface{}]interface{}), callback)

	default:
		jBytes, err := marshal(v)
		if err != nil {

			return err
		}

		callback(jBytes)

		return nil
	}
}

func readArrayByString(v string, callback func([]byte)) error {
	callback([]byte(v))

	return nil
}

func readArrayBySliceString(v []string, callback func([]byte)) error {
	for i := range v {
		callback([]byte(v[i]))
	}

	return nil
}

func readArrayBySliceInterface(marshal func(interface{}) ([]byte, error), v []interface{}, callback func([]byte)) error {
	if len(v) == 0 {
		return nil
	}

	switch v[0].(type) {
	case string:
		for i := range v {
			callback([]byte(v[i].(string)))
		}

	case []byte:
		for i := range v {
			callback(v[i].([]byte))
		}

	default:
		for i := range v {

			jBytes, err := marshal(v[i])
			if err != nil {
				return err
			}
			callback(jBytes)
		}
	}

	return nil
}

func readArrayByMapIfaceIface(marshal func(interface{}) ([]byte, error), v map[interface{}]interface{}, callback func([]byte)) error {
	for key, val := range v {

		bKey := []byte(fmt.Sprint(key) + ": ")
		b, err := marshal(val)
		if err != nil {
			return err
		}
		callback(append(bKey, b...))
	}

	return nil
}

func readArrayByMapStrStr(v map[string]string, callback func([]byte)) error {
	for key, val := range v {

		callback([]byte(key + ": " + val))
	}

	return nil
}

func readArrayByMapStrIface(marshal func(interface{}) ([]byte, error), v map[string]interface{}, callback func([]byte)) error {
	for key, val := range v {

		bKey := []byte(key + ": ")
		b, err := marshal(val)
		if err != nil {
			return err
		}
		callback(append(bKey, b...))
	}

	return nil
}

func readArrayByMapIfaceStr(v map[interface{}]string, callback func([]byte)) error {
	for key, val := range v {

		callback([]byte(fmt.Sprint(key) + ": " + val))
	}

	return nil
}
