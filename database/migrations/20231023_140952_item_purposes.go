package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ItemPurposes_20231023_140952 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ItemPurposes_20231023_140952{}
	m.Created = "20231023_140952"

	migration.Register("ItemPurposes_20231023_140952", m)
}

// Run the migrations
func (m *ItemPurposes_20231023_140952) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE item_purposes(`item_purpose_id` int(11) NOT NULL AUTO_INCREMENT,`item_id` int(11) DEFAULT NULL,`purpose_id` int(11) DEFAULT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,PRIMARY KEY (`item_purpose_id`), FOREIGN KEY (item_id) REFERENCES items(item_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (purpose_id) REFERENCES purposes(purpose_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *ItemPurposes_20231023_140952) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `item_purposes`")
}
