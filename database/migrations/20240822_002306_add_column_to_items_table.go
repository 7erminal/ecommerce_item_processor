package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToItemsTable_20240822_002306 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToItemsTable_20240822_002306{}
	m.Created = "20240822_002306"

	migration.Register("AddColumnToItemsTable_20240822_002306", m)
}

// Run the migrations
func (m *AddColumnToItemsTable_20240822_002306) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE items add COLUMN material varchar(400) default '0'")
	m.SQL("ALTER TABLE items add COLUMN country int(11) default 1, ADD FOREIGN KEY (country) REFERENCES countries(country_id) ON UPDATE CASCADE ON DELETE NO ACTION")
}

// Reverse the migrations
func (m *AddColumnToItemsTable_20240822_002306) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
