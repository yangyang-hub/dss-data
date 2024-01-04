package util

import (
	"bytes"
	"reflect"
)

//根据结构体拼接cypher
func BuildNodeProperties(node interface{}, mark, prefix string) (string, map[string]interface{}) {
	buffer := bytes.Buffer{}
	param := make(map[string]interface{})
	n := reflect.TypeOf(node)
	for i := 0; i < n.NumField(); i++ {
		field := n.Field(i)
		fieldName := field.Name
		jsonName := field.Tag.Get("json")
		fieldType := field.Type.Name()
		switch fieldType {
		case "string":
			param[jsonName] = reflect.ValueOf(node).FieldByName(fieldName).String()
		case "int64":
			param[jsonName] = reflect.ValueOf(node).FieldByName(fieldName).Int()
		case "float64":
			param[jsonName] = reflect.ValueOf(node).FieldByName(fieldName).Float()
		}
		buffer.WriteString(prefix)
		buffer.WriteString(jsonName)
		buffer.WriteString(mark)
		buffer.WriteString(" $")
		buffer.WriteString(jsonName)

		if i < n.NumField()-1 {
			buffer.WriteString(",")
		}
	}
	return buffer.String(), param
}
