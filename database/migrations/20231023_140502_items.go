package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Items_20231023_140502 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Items_20231023_140502{}
	m.Created = "20231023_140502"

	migration.Register("Items_20231023_140502", m)
}

// Run the migrations
func (m *Items_20231023_140502) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE items(`item_id` int(11) NOT NULL AUTO_INCREMENT,`item_name` varchar(80) NOT NULL,`description` varchar(250) DEFAULT NULL,`category` int(11) NOT NULL,`available_sizes` varchar(250) DEFAULT NULL,`available_colors` varchar(250) DEFAULT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,PRIMARY KEY (`item_id`), FOREIGN KEY (category) REFERENCES categories(category_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *Items_20231023_140502) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `items`")
}
