package functions

import (
	"bytes"
	"encoding/json"
	"io"
	"item_processor/api"
	"item_processor/structs/requests"
	"item_processor/structs/responses"
	"strconv"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func setControllerJSON(c *beego.Controller, value interface{}) {
	if c == nil {
		return
	}
	if c.Data == nil {
		c.Data = map[interface{}]interface{}{}
	}
	c.Data["json"] = value
}

func AddBranch(c *beego.Controller, req requests.BranchRequestDTO, addedBy int64) (resp responses.BranchOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Sending user name ", strconv.FormatInt(addedBy, 10))

	request := api.NewRequest(
		host,
		"/v1/branches/",
		api.POST)
	request.InterfaceParams["Branch"] = req.Branch
	request.InterfaceParams["CountryCode"] = req.CountryCode
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["Location"] = req.Location
	request.InterfaceParams["AddedBy"] = strconv.FormatInt(addedBy, 10)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.BranchOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func UpdateBranch(c *beego.Controller, req requests.BranchRequestDTO, addedBy int64, branchId string) (resp responses.BranchOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Sending user name ", strconv.FormatInt(addedBy, 10))

	request := api.NewRequest(
		host,
		"/v1/branches/"+branchId,
		api.PUT)
	request.InterfaceParams["Branch"] = req.Branch
	request.InterfaceParams["CountryCode"] = req.CountryCode
	request.InterfaceParams["PhoneNumber"] = req.PhoneNumber
	request.InterfaceParams["Location"] = req.Location
	request.InterfaceParams["AddedBy"] = strconv.FormatInt(addedBy, 10)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.BranchOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetBranch(c *beego.Controller, branchid int64) (resp responses.BranchOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Getting branch details for ", branchid)

	request := api.NewRequest(
		host,
		"/v1/branches/"+strconv.FormatInt(branchid, 10),
		api.GET)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.BranchOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func DeleteBranch(c *beego.Controller, branchid string) (resp responses.StringOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Getting branch details for ", branchid)

	request := api.NewRequest(
		host,
		"/v1/branches/"+branchid,
		api.DELETE)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.StringOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetBranches(c *beego.Controller) (resp responses.BranchesOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/branches/",
		api.GET)
	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.BranchesOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func UpdateBranchBranchManger(c *beego.Controller, userid string, branchid string) (resp responses.BranchesOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/branches/branch-manager/"+branchid,
		api.PUT)
	request.InterfaceParams["BranchManager"] = userid
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.BranchesOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCountries(c *beego.Controller) (resp responses.CountriesOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/countries/",
		api.GET)

	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.CountriesOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCountry(c *beego.Controller, countryId int64) (resp responses.CountriesOriResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/countries/"+strconv.FormatInt(countryId, 10),
		api.GET)

	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.CountriesOriResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCountryWithCode(c *beego.Controller, countryCode string) (resp responses.CountryResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/countries/code/"+countryCode,
		api.GET)

	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.CountryResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCurrencies(c *beego.Controller) (resp responses.CurrenciesResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/currencies/",
		api.GET)

	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.CurrenciesResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCurrency(c *beego.Controller, currencyId int64) (resp responses.CurrenciesResponseDTO) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	request := api.NewRequest(
		host,
		"/v1/currencies/"+strconv.FormatInt(currencyId, 10),
		api.GET)

	// request.Params["Dob"] = req.Dob
	// request.Params["Gender"] = req.Gender
	// request.Params["PhoneNumber"] = req.PhoneNumber
	// request.Params["Username"] = req.Username
	// request.Params["MaritalStatus"] = ""
	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	logs.Info("Raw response received is ", res)
	// data := map[string]interface{}{}
	var data responses.CurrenciesResponseDTO
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)

	return data
}

func GetCurrencyWithName(c *beego.Controller, currencyName string) (responses.CurrencyResponseDTO, error) {
	host, _ := beego.AppConfig.String("systemBaseUrl")

	logs.Info("Request to get Currency: ", currencyName)

	request := api.NewRequest(
		host,
		"/v1/currencies/"+currencyName,
		api.GET)

	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "params",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		setControllerJSON(c, err.Error())
		return responses.CurrencyResponseDTO{}, err
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		setControllerJSON(c, err.Error())
		return responses.CurrencyResponseDTO{}, err
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	// var dataOri responses.UserOriResponseDTO
	var data responses.CurrencyResponseDTO
	if err := json.Unmarshal(read, &data); err != nil {
		setControllerJSON(c, err.Error())
		return responses.CurrencyResponseDTO{}, err
	}
	setControllerJSON(c, data)

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data, nil
}
