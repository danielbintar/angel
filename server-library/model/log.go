package model

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/danielbintar/angel/server-library/pubsub"
)

func Log(micro string, model interface{}, pub pubsub.AsyncPublisher) {
	changes := GenerateChanges(model)

	if len(changes) == 0 { return }

	payload := RequestPayload {
		ID: fmt.Sprintf("%v", reflect.ValueOf(model).FieldByName("ID").Interface()),
		ModelName: reflect.TypeOf(model).Name(),
		MicroName: micro,
		Changes: changes,
	}

	encodedPayload, _ := json.Marshal(payload)

	pub.Publish(micro + "_model-log", string(encodedPayload))
}

type RequestPayload struct {
	ID        string   `json:"id"`
	ModelName string   `json:"model_name"`
	MicroName string   `json:"micro_name"`
	Changes   []Change `json:"changes"`
}

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

		changes = append(changes, Change {
			Key: k,
			After: afterVal,
			Previous: prevVal,
		})
	}

	return changes
}
