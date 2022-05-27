package appstoreserverapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// ApiSendConsumptionInformation 发送消费信息
// Send Consumption Information
// doc: https://developer.apple.com/documentation/appstoreserverapi/send_consumption_information
func (c *client) ApiSendConsumptionInformation(transactionId string, req ConsumptionRequest) error {
	reqUrl := c.api_send_consumption_information_url + transactionId
	b, _ := json.Marshal(req)
	_, err := c.doRequest(http.MethodPut, reqUrl, bytes.NewReader(b))
	if err != nil {
		return err
	}
	return nil
}

// ConsumptionRequest 消费请求
// doc: https://developer.apple.com/documentation/appstoreserverapi/consumptionrequest
type ConsumptionRequest struct {
	AccountTenure            uint8  `json:"accountTenure"`
	AppAccountToken          string `json:"appAccountToken"`
	ConsumptionStatus        uint8  `json:"consumptionStatus"`
	CustomerConsented        bool   `json:"customerConsented"`
	DeliveryStatus           uint8  `json:"deliveryStatus"`
	LifetimeDollarsPurchased uint8  `json:"lifetimeDollarsPurchased"`
	LifetimeDollarsRefunded  uint8  `json:"lifetimeDollarsRefunded"`
	Platform                 uint8  `json:"platform"`
	PlayTime                 uint8  `json:"playTime"`
	SampleContentProvided    bool   `json:"sampleContentProvided"`
	UserStatus               uint8  `json:"userStatus"`
}
