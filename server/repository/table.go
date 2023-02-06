package repository

import (
	"fmt"
)

type Table struct {
	Name   string
	Alias  string
	Fields []string
}

func NewTable(name string, alias string, fields []string) *Table {
	return &Table{
		Name:   name,
		Alias:  alias,
		Fields: fields,
	}
}

func (t *Table) NameWithAlias() string {
	return NewArgs(t.Name, t.Alias).ToString()
}

func (t *Table) FieldWithAlias(fieldName string) string {
	return NewArgs(t.Alias, fieldName).ToString(".")
}

func (t *Table) FieldsWithAlias() string {
	args := NewArgs()
	for _, fieldName := range t.Fields {
		args = args.Append(t.FieldWithAlias(fieldName))
	}
	return args.ToString(", ")
}

func (t *Table) Select(args Args) string {
	return fmt.Sprintf("SELECT %s FROM %s %s", t.FieldsWithAlias(), t.NameWithAlias(), args.ToString())
}

func (t *Table) SelectDistinct(args Args) string {
	return fmt.Sprintf("SELECT DISTINCT %s FROM %s %s", t.FieldsWithAlias(), t.Name, args.ToString())
}
