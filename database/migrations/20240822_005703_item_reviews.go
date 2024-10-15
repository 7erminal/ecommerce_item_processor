package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ItemReviews_20240822_005703 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ItemReviews_20240822_005703{}
	m.Created = "20240822_005703"

	migration.Register("ItemReviews_20240822_005703", m)
}

// Run the migrations
func (m *ItemReviews_20240822_005703) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE item_reviews(`item_review_id` int(11) NOT NULL AUTO_INCREMENT,`review` varchar(1000) DEFAULT NULL,`rating` int(11) DEFAULT NULL,`review_by` int(11) DEFAULT NULL,`date_created` datetime DEFAULT CURRENT_TIMESTAMP,`date_modified` datetime ON UPDATE CURRENT_TIMESTAMP,`created_by` int(11) DEFAULT NULL,`modified_by` int(11) DEFAULT NULL,`active` int(11) DEFAULT 1,PRIMARY KEY (`item_review_id`), FOREIGN KEY (review_by) REFERENCES users(user_id) ON UPDATE CASCADE ON DELETE NO ACTION)")
}

// Reverse the migrations
func (m *ItemReviews_20240822_005703) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `item_reviews`")
}
