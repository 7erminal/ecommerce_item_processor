package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Purposes_20231023_135440 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Purposes_20231023_135440{}
	m.Created = "20231023_135440"

	migration.Register("Purposes_20231023_135440", m)
}

// Run the migrations
func (m *Purposes_20231023_135440) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE purposes(`purpose_id` int(11) NOT NULL AUTO_INCREMENT,`purpose` varchar(40) NOT NULL,`image_path` varchar(250) NOT NULL,`description` varchar(250) DEFAULT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,PRIMARY KEY (`purpose_id`))")
}

// Reverse the migrations
func (m *Purposes_20231023_135440) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `purposes`")
}
