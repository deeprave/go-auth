package models

import (
	"time"
)

type Field struct {
	Name   string `json:"name,omitempty"`
	Type   string `json:"type"`
	Length int    `json:"length,omitempty"`
}

type Table struct {
	Oid          int64     `json:"oid"`
	Name         string    `json:"name"`
	Fields       []Field   `json:"fields"`
	NumRecords   int64     `json:"num_records"`
	NumAvailable int64     `json:"num_available"`
	FirstId      int64     `json:"first_id,omitempty"`
	LastId       int64     `json:"last_id,omitempty"`
	FirstCreated time.Time `json:"first_created"`
	LastCreated  time.Time `json:"last_created"`
	LastUpdated  time.Time `json:"last_updated"`
}

func (table *Table) TableName() string {
	return "\"" + table.Name + "\""
}

func (table *Table) FindFieldByName(name string) (int, *Field) {
	for i, field := range table.Fields {
		if field.Name == name {
			return i, &field
		}
	}
	return -1, nil
}

func (table *Table) HasFieldName(name string) bool {
	i, _ := table.FindFieldByName(name)
	return i != -1
}

type Schema struct {
	Oid       int64    `json:"oid"`
	Name      string   `json:"name"`
	Sequences []string `json:"sequences,omitempty"`
	Tables    []Table  `json:"tables,omitempty"`
}

type Database struct {
	Database  string   `json:"database"`
	Commits   int64    `json:"commits"`
	Rollbacks int64    `json:"rollbacks"`
	Sessions  int64    `json:"sessions"`
	Deadlocks int64    `json:"deadlocks"`
	Schemas   []Schema `json:"schemas,omitempty"`
}
