package functions

import (
	"item_processor/models"
	"item_processor/structs/requests"

	"github.com/beego/beego/v2/core/logs"
)

func CheckItemCount(itemid int64, itemName string) (status bool) {
	status = false
	if quantity, err := models.GetItem_quantityByItemId(itemid); err == nil {
		if quantity.Quantity <= quantity.QuantityAlert {
			// serviceName := "ORDER"
			// status := "WARNING"
			// category := "Stock"

			// quantityLeft := strconv.Itoa(quantity.Quantity)

			// SendNotification(serviceName, status, category, itemName, quantityLeft, nil)
			return true

		}
	}

	return status
}

func SendNotification(serviceName string, status string, category string, itemName string, quantityLeft string, userid *int64) {
	notification := requests.NotificationRequest{
		UserId:   userid,
		Service:  serviceName,
		Status:   status,
		Category: category,
		Params:   []string{itemName, quantityLeft},
	}

	resp := SendNotificationApi(notification)

	if resp.StatusCode == 200 {
		logs.Info("Successfully added notification")
	} else {
		logs.Info("Notification addition failed")
	}
}
