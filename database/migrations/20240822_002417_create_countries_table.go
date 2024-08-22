package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateCountriesTable_20240822_002417 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateCountriesTable_20240822_002417{}
	m.Created = "20240822_002417"

	migration.Register("CreateCountriesTable_20240822_002417", m)
}

// Run the migrations
func (m *CreateCountriesTable_20240822_002417) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE countries(`country_id` int(11) NOT NULL AUTO_INCREMENT,`country` varchar(80) NOT NULL,`description` varchar(250) DEFAULT NULL,`country_code` varchar(250) DEFAULT NULL,`default_currency` int(11) DEFAULT 1,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,PRIMARY KEY (`country_id`), FOREIGN KEY (default_currency) REFERENCES currencies(currency_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *CreateCountriesTable_20240822_002417) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
