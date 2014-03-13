package puck

import (
	"bytes"
	"fmt"
)

type Resources struct {
	Resources []Resource
}

type Resource struct {
	Type          string
	NameAttribute string `json:"name_attribute"`
	Attributes    map[string]interface{}
}

func (r *Resource) ToChefResource() string {
	var buffer bytes.Buffer
	quotedNameAttribute := fmt.Sprintf("\"%v\"", r.NameAttribute)
	buffer.WriteString(fmt.Sprintln(r.Type, quotedNameAttribute, "do"))

	for k, v := range r.Attributes {
		switch v := v.(type) {
		case string:
			quotedString := fmt.Sprintf("\"%v\"", v)
			buffer.WriteString(fmt.Sprintln("  ", k, quotedString))
		default:
			buffer.WriteString(fmt.Sprintln("  ", k, v))
		}
	}

	buffer.WriteString(fmt.Sprintln("end"))
	return buffer.String()
}
