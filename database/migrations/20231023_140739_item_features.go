package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ItemFeatures_20231023_140739 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ItemFeatures_20231023_140739{}
	m.Created = "20231023_140739"

	migration.Register("ItemFeatures_20231023_140739", m)
}

// Run the migrations
func (m *ItemFeatures_20231023_140739) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE item_features(`item_feature_id` int(11) NOT NULL AUTO_INCREMENT,`item_id` int(11) NOT NULL,`feature_id` int(11) NOT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,PRIMARY KEY (`item_feature_id`), FOREIGN KEY (item_id) REFERENCES items(item_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (feature_id) REFERENCES features(feature_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *ItemFeatures_20231023_140739) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `item_features`")
}
