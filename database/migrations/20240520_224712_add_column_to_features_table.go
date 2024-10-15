package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToFeaturesTable_20240520_224712 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToFeaturesTable_20240520_224712{}
	m.Created = "20240520_224712"

	migration.Register("AddColumnToFeaturesTable_20240520_224712", m)
}

// Run the migrations
func (m *AddColumnToFeaturesTable_20240520_224712) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE features add COLUMN visible bool default 0")
}

// Reverse the migrations
func (m *AddColumnToFeaturesTable_20240520_224712) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
