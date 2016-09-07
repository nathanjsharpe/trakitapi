package trakitapi

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func getStrValues(v reflect.Value) []string {
	values := make([]string, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		val, err := getStrValue(v.Field(i).Interface())
		if err != nil {
			fmt.Println(err.Error(), v.Type().Field(i).Name)
			os.Exit(1)
		}
		values[i] = val
	}
	return values
}

func getStrValue(val interface{}) (string, error) {
	switch val.(type) {
	case string:
		return val.(string), nil
	case bool:
		return strconv.FormatBool(val.(bool)), nil
	case int:
		return strconv.Itoa(val.(int)), nil
	case float64:
		return strconv.FormatFloat(val.(float64), 'f', 5, 64), nil
	default:
		return "", errors.New("Uknown conversion for value")
	}
}
