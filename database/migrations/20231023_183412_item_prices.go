package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ItemPrices_20231023_183412 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ItemPrices_20231023_183412{}
	m.Created = "20231023_183412"

	migration.Register("ItemPrices_20231023_183412", m)
}

// Run the migrations
func (m *ItemPrices_20231023_183412) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE item_prices(`item_price_id` int(11) NOT NULL AUTO_INCREMENT,`item_id` int(11) NOT NULL,`item_price` int(11) DEFAULT 0,`currency_id` int(11) NOT NULL,`active` int(11) DEFAULT 1,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT 1,`modified_by` int(11) DEFAULT 1,PRIMARY KEY (`item_price_id`), FOREIGN KEY (item_id) REFERENCES items(item_id) ON UPDATE CASCADE ON DELETE NO ACTION, FOREIGN KEY (currency_id) REFERENCES currencies(currency_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *ItemPrices_20231023_183412) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `item_prices`")
}
