package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToItemsTable_20240722_234818 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToItemsTable_20240722_234818{}
	m.Created = "20240722_234818"

	migration.Register("AddColumnToItemsTable_20240722_234818", m)
}

// Run the migrations
func (m *AddColumnToItemsTable_20240722_234818) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE items add COLUMN weight varchar(20) default '0'")
}

// Reverse the migrations
func (m *AddColumnToItemsTable_20240722_234818) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
