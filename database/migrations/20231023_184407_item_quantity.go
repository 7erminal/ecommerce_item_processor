package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ItemQuantity_20231023_184407 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ItemQuantity_20231023_184407{}
	m.Created = "20231023_184407"

	migration.Register("ItemQuantity_20231023_184407", m)
}

// Run the migrations
func (m *ItemQuantity_20231023_184407) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE item_quantity(`item_quantity_id` int(11) NOT NULL AUTO_INCREMENT,`item_id` int(11) NOT NULL,`quantity` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,PRIMARY KEY (`item_quantity_id`), FOREIGN KEY (item_id) REFERENCES items(item_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *ItemQuantity_20231023_184407) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `item_quantity`")
}
