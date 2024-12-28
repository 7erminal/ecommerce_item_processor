package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CreateBranchesTable_20241227_115307 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CreateBranchesTable_20241227_115307{}
	m.Created = "20241227_115307"

	migration.Register("CreateBranchesTable_20241227_115307", m)
}

// Run the migrations
func (m *CreateBranchesTable_20241227_115307) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE branches(`branch_id` int(11) NOT NULL AUTO_INCREMENT,`branch` varchar(80) NOT NULL,`country` int(11) DEFAULT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,PRIMARY KEY (`branch_id`), FOREIGN KEY (country) REFERENCES countries(country_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *CreateBranchesTable_20241227_115307) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
