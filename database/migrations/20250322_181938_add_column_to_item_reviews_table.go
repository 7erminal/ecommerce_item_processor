package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToItemReviewsTable_20250322_181938 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToItemReviewsTable_20250322_181938{}
	m.Created = "20250322_181938"

	migration.Register("AddColumnToItemReviewsTable_20250322_181938", m)
}

// Run the migrations
func (m *AddColumnToItemReviewsTable_20250322_181938) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE item_reviews ADD COLUMN item_id int DEFAULT NULL, ADD COLUMN reference int DEFAULT NULL, ADD FOREIGN KEY (item_id) REFERENCES items(item_id) ON DELETE NO ACTION ON UPDATE NO ACTION")

}

// Reverse the migrations
func (m *AddColumnToItemReviewsTable_20250322_181938) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
