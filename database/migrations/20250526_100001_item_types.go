package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ItemTypes_20250526_100001 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ItemTypes_20250526_100001{}
	m.Created = "20250526_100001"

	migration.Register("ItemTypes_20250526_100001", m)
}

// Run the migrations
func (m *ItemTypes_20250526_100001) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE item_types(`item_type_id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(80) NOT NULL,`description` varchar(255) NOT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,PRIMARY KEY (`item_type_id`))")
}

// Reverse the migrations
func (m *ItemTypes_20250526_100001) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `item_types`")
}
