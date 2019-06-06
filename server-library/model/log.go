package model

import (
	"fmt"
	"reflect"
)

type Change struct {
	Key      string
	Previous string
	After    string
}

func GenerateChanges(model interface{}) []Change {
	var changes []Change

	t := reflect.TypeOf(model)
	v := reflect.ValueOf(model)

	prev := v.FieldByName("PreviousData").Elem()

	for i := 0; i < v.NumField(); i++ {
		k := t.Field(i).Name
		if k == "PreviousData" || k == "ID" { continue }

		prevVal := fmt.Sprintf("%v", prev.Field(i).Interface())
		afterVal := fmt.Sprintf("%v", v.Field(i).Interface())

		if prevVal == afterVal { continue }

		change := Change {
			Key: k,
			After: afterVal,
			Previous: prevVal,
		}
		changes = append(changes, change)
	}

	return changes
}
