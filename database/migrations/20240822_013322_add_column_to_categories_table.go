package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToCategoriesTable_20240822_013322 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToCategoriesTable_20240822_013322{}
	m.Created = "20240822_013322"

	migration.Register("AddColumnToCategoriesTable_20240822_013322", m)
}

// Run the migrations
func (m *AddColumnToCategoriesTable_20240822_013322) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE categories add COLUMN icon varchar(400) default NULL")
}

// Reverse the migrations
func (m *AddColumnToCategoriesTable_20240822_013322) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
