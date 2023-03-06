package lib

import (
	"fmt"
)

type TableDef struct {
	Name   string
	Alias  string
	Fields []string
}

func NewTable(name string, alias string, fields []string) *TableDef {
	return &TableDef{
		Name:   name,
		Alias:  alias,
		Fields: fields,
	}
}

func (t *TableDef) NameWithAlias() string {
	return NewArgs(t.Name, t.Alias).ToString()
}

func (t *TableDef) FieldWithAlias(fieldName string) string {
	return NewArgs(t.Alias, fieldName).ToString(".")
}

func (t *TableDef) FieldsWithAlias() string {
	args := NewArgs()
	for _, fieldName := range t.Fields {
		args = args.Append(t.FieldWithAlias(fieldName))
	}
	return args.ToString(", ")
}

func (t *TableDef) Select(args Args) string {
	return fmt.Sprintf("SELECT %s FROM %s %s", t.FieldsWithAlias(), t.NameWithAlias(), args.ToString())
}

func (t *TableDef) SelectDistinct(args Args) string {
	return fmt.Sprintf("SELECT DISTINCT %s FROM %s %s", t.FieldsWithAlias(), t.Name, args.ToString())
}
