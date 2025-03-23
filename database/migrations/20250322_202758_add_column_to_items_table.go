package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToItemsTable_20250322_202758 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToItemsTable_20250322_202758{}
	m.Created = "20250322_202758"

	migration.Register("AddColumnToItemsTable_20250322_202758", m)
}

// Run the migrations
func (m *AddColumnToItemsTable_20250322_202758) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE items ADD COLUMN item_status int DEFAULT NULL, ADD FOREIGN KEY (status_id) REFERENCES status(status_id) ON DELETE NO ACTION ON UPDATE NO ACTION")
}

// Reverse the migrations
func (m *AddColumnToItemsTable_20250322_202758) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
