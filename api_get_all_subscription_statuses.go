package appstoreserverapi

import (
	"net/http"
)

// ApiGetAllSubscriptionStatuses 获取所有的订阅状态
// Get All Subscription Statuses
// doc: https://developer.apple.com/documentation/appstoreserverapi/get_all_subscription_statuses
func (c *client) ApiGetAllSubscriptionStatuses(transactionId string) (*StatusResponse, error) {
	reqUrl := c.apiGetAllSubscriptionStatusesUrl + transactionId
	r, err := c.doRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}

	// decode
	result := &StatusResponse{
		raw:         r.String(),
		Environment: r.Get("environment").String(),
		BundleId:    r.Get("bundleId").String(),
		AppAppleId:  r.Get("appAppleId").Int(),
	}

	datas := make([]StatusData, 0)

	for _, item := range r.Get("data").Array() {
		data := StatusData{
			SubscriptionGroupIdentifier: item.Get("subscriptionGroupIdentifier").String(),
		}
		lastTransactions := make([]LastTransaction, 0)
		for _, val := range item.Get("lastTransactions").Array() {
			lastTransaction := LastTransaction{
				OriginalTransactionId: val.Get("originalTransactionId").String(),
				Status:                val.Get("status").Int(),
			}
			jWSTransactionDecodedPayload := JWSTransactionDecodedPayload{}
			Parse(val.Get("signedTransactionInfo").String(), &jWSTransactionDecodedPayload)
			lastTransaction.SignedTransactionInfo = jWSTransactionDecodedPayload

			jWSRenewalInfoDecodedPayload := JWSRenewalInfoDecodedPayload{}
			Parse(val.Get("signedRenewalInfo").String(), &jWSRenewalInfoDecodedPayload)
			lastTransaction.SignedRenewalInfo = jWSRenewalInfoDecodedPayload

			lastTransactions = append(lastTransactions, lastTransaction)
		}
		data.LastTransactions = lastTransactions
		datas = append(datas, data)
	}

	result.Data = datas

	return result, nil
}

type StatusResponse struct {
	raw         string
	Environment string       `json:"environment"`
	BundleId    string       `json:"bundleId"`
	AppAppleId  int64        `json:"appAppleId"`
	Data        []StatusData `json:"data"`
}

type StatusData struct {
	SubscriptionGroupIdentifier string            `json:"subscriptionGroupIdentifier"`
	LastTransactions            []LastTransaction `json:"lastTransactions"`
}

type LastTransaction struct {
	OriginalTransactionId string                       `json:"originalTransactionId"`
	Status                int64                        `json:"status"`
	SignedTransactionInfo JWSTransactionDecodedPayload `json:"signedTransactionInfo"`
	SignedRenewalInfo     JWSRenewalInfoDecodedPayload `json:"signedRenewalInfo"`
}

// Raw 返回原始结果
func (r *StatusResponse) Raw() string {
	return r.raw
}
