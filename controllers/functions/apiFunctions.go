package functions

import (
	"encoding/json"
	"io"
	"item_processor/api"
	"item_processor/structs/requests"
	"item_processor/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func SendNotificationApi(req requests.NotificationRequest) (resp responses.NotificationResponse) {
	host, _ := beego.AppConfig.String("notificationsBaseUrl")

	// logs.Info("Sending first name ", req.BranchId)

	request := api.NewRequest(
		host,
		"/v1/notifications/",
		api.POST)
	request.InterfaceParams["UserId"] = req.UserId
	request.InterfaceParams["Category"] = req.Category
	request.InterfaceParams["Service"] = req.Service
	request.InterfaceParams["Status"] = req.Status
	request.InterfaceParams["Params"] = req.Params

	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		// c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		logs.Error("Error::: ", err.Error())
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.NotificationResponse
	json.Unmarshal(read, &data)

	logs.Info("Resp is ", data)

	return data
}
