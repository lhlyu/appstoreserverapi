package appstoreserverapi

import (
	"net/http"
	"sort"
	"strings"
)

// ApiGetTransactionHistory 获取历史交易记录
// Get Transaction History
// doc: https://developer.apple.com/documentation/appstoreserverapi/get_transaction_history
// desc: true then signedTransactions order by webOrderLineItemId desc
func (c *client) ApiGetTransactionHistory(transactionId string, desc bool) (*HistoryResponse, error) {
	reqUrl := c.apiGetTransactionHistoryUrl + transactionId
	r, err := c.doRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	result := &HistoryResponse{
		raw:         r.String(),
		Revision:    r.Get("revision").String(),
		BundleId:    r.Get("bundleId").String(),
		AppAppleId:  r.Get("appAppleId").Int(),
		Environment: r.Get("environment").String(),
		HasMore:     r.Get("hasMore").Bool(),
	}

	signedTransactions := make([]JWSTransactionDecodedPayload, 0)
	for _, item := range r.Get("signedTransactions").Array() {
		signedTransaction := JWSTransactionDecodedPayload{}
		Parse(item.String(), &signedTransaction)
		signedTransactions = append(signedTransactions, signedTransaction)
	}

	if desc {
		sort.SliceStable(signedTransactions, func(i, j int) bool {
			if strings.Compare(signedTransactions[i].WebOrderLineItemId, signedTransactions[j].WebOrderLineItemId) == 1 {
				return true
			}
			return false
		})
	}

	result.SignedTransactions = signedTransactions

	return result, nil
}

type HistoryResponse struct {
	raw                string
	Revision           string                         `json:"revision"`
	BundleId           string                         `json:"bundleId"`
	AppAppleId         int64                          `json:"appAppleId"`
	Environment        string                         `json:"environment"`
	HasMore            bool                           `json:"hasMore"`
	SignedTransactions []JWSTransactionDecodedPayload `json:"signedTransactions"`
}

func (r *HistoryResponse) Raw() string {
	return r.raw
}
