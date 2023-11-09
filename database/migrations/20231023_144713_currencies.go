package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Currencies_20231023_144713 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Currencies_20231023_144713{}
	m.Created = "20231023_144713"

	migration.Register("Currencies_20231023_144713", m)
}

// Run the migrations
func (m *Currencies_20231023_144713) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE currencies(`currency_id` int(11) NOT NULL AUTO_INCREMENT,`symbol` varchar(20) NOT NULL,`currency` varchar(50) NOT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,PRIMARY KEY (`currency_id`))")
}

// Reverse the migrations
func (m *Currencies_20231023_144713) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `currencies`")
}
