package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type AddColumnToItemPriceTable_20240822_000708 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddColumnToItemPriceTable_20240822_000708{}
	m.Created = "20240822_000708"

	migration.Register("AddColumnToItemPriceTable_20240822_000708", m)
}

// Run the migrations
func (m *AddColumnToItemPriceTable_20240822_000708) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE item_prices add COLUMN discount varchar(20) default '0'")
	m.SQL("ALTER TABLE item_prices add COLUMN discount_type varchar(20) default 'PERCENTAGE'")
}

// Reverse the migrations
func (m *AddColumnToItemPriceTable_20240822_000708) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
