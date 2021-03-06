package schema

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Type defines a specific type within the schema
type Type struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Kind        Kind   `json:"kind,omitempty"`

	EnumValues    []EnumValue `json:"enumValues,omitempty"`
	Fields        []Field     `json:"fields,omitempty"`
	InputFields   []Field     `json:"inputFields,omitempty"`
	Interfaces    []TypeRef   `json:"interfaces,omitempty"`
	PossibleTypes []TypeRef   `json:"possibleTypes,omitempty"`
}

// Save writes the schema out to a file
func (t *Type) Save(file string) error {
	if file == "" {
		return errors.New("unable to save schema, no file specified")
	}

	log.WithFields(log.Fields{
		"schema_file": file,
	}).Debug("saving schema")

	schemaFile, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, schemaFile, 0644)
}

// GetDescription formats the description into a GoDoc comment.
func (t *Type) GetDescription() string {
	if strings.TrimSpace(t.Description) == "" {
		return ""
	}

	return formatDescription(t.GetName(), t.Description)
}

// GetName returns the name of a Type, formatted for Go title casing.
func (t *Type) GetName() string {
	return formatGoName(t.Name)
}

// IsGoType is used to determine if a type in NerdGraph is already a native type of Golang.
func (t *Type) IsGoType() bool {
	goTypes := []string{
		"int",
		"string",
		"bool",
		"boolean",
	}

	name := strings.ToLower(t.GetName())

	for _, x := range goTypes {
		if x == name {
			return true
		}
	}

	return false
}
