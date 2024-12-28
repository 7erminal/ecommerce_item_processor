package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToItemsTable_20241227_120607 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToItemsTable_20241227_120607{}
	m.Created = "20241227_120607"

	migration.Register("AddColumnToItemsTable_20241227_120607", m)
}

// Run the migrations
func (m *AddColumnToItemsTable_20241227_120607) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE items add COLUMN branch int(11) default NULL, ADD FOREIGN KEY (branch) REFERENCES branches(branch_id) ON UPDATE CASCADE ON DELETE NO ACTION")
}

// Reverse the migrations
func (m *AddColumnToItemsTable_20241227_120607) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
