package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToPurposesTable_20240520_224721 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToPurposesTable_20240520_224721{}
	m.Created = "20240520_224721"

	migration.Register("AddColumnToPurposesTable_20240520_224721", m)
}

// Run the migrations
func (m *AddColumnToPurposesTable_20240520_224721) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE purposes add COLUMN visible bool default 0")
}

// Reverse the migrations
func (m *AddColumnToPurposesTable_20240520_224721) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
