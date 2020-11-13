package db

import (
	"errors"
	"net/url"
	"reflect"
)

type QueryFilter map[string]interface{}

func (qf QueryFilter) GetFromMap(data interface{}, Allowed []string) error {
	rData := reflect.ValueOf(data)

	if rData.Kind() != reflect.Map {
		return errors.New("Input is not map type")
	}

	for _, key := range rData.MapKeys() {
		keyName := key.Interface()

		// check if key is in Allowed
		if Allowed == nil {
			goto ProcessKey
		}

		for _, aKey := range Allowed {
			if aKey == keyName {
				goto ProcessKey
			}
		}
		continue

	ProcessKey:

		qf[keyName.(string)] = rData.MapIndex(key)
	}

	return nil
}

func (qf QueryFilter) GetFromURLQuery(uq url.Values, Allowed []string) error {
	if Allowed != nil {
		for _, k := range Allowed {
			if v, ok := uq[k]; ok {
				if len(v) == 1 {
					qf[k] = v[0]
				} else {
					qf[k] = v
				}
			}
		}

		return nil
	}

	for k, v := range uq {
		if len(v) == 1 {
			qf[k] = v[0]
		} else {
			qf[k] = v
		}
	}

	return nil
}
