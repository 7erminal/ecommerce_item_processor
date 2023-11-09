package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ItemImages_20231023_144125 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ItemImages_20231023_144125{}
	m.Created = "20231023_144125"

	migration.Register("ItemImages_20231023_144125", m)
}

// Run the migrations
func (m *ItemImages_20231023_144125) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE item_images(`item_image_id` int(11) NOT NULL AUTO_INCREMENT,`item_id` int(11) NOT NULL,`image_path` varchar(250) NOT NULL,`is_default` int(11) DEFAULT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,PRIMARY KEY (`item_image_id`), FOREIGN KEY (item_id) REFERENCES items(item_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *ItemImages_20231023_144125) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `item_images`")
}
