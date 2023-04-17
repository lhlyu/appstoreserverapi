package appstoreserverapi

import (
	"net/http"
)

// ApiLookUpOrderId 查找订单 ID
// Look Up Order ID
// doc: https://developer.apple.com/documentation/appstoreserverapi/look_up_order_id
func (c *client) ApiLookUpOrderId(orderId string) (*OrderLookupResponse, error) {
	reqUrl := c.apiLookupOrderIdUrl + orderId
	r, err := c.doRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	result := &OrderLookupResponse{
		raw:    r.String(),
		Status: r.Get("status").Int(),
	}

	signedTransactions := make([]JWSTransactionDecodedPayload, 0)
	for _, item := range r.Get("signedTransactions").Array() {
		signedTransaction := JWSTransactionDecodedPayload{}
		Parse(item.String(), &signedTransaction)
		signedTransactions = append(signedTransactions, signedTransaction)
	}
	result.SignedTransactions = signedTransactions

	return result, nil
}

type OrderLookupResponse struct {
	raw    string
	Status int64 `json:"status"`
	// 原始数据
	SignedTransactions []JWSTransactionDecodedPayload `json:"signedTransactions"`
}

func (r *OrderLookupResponse) Raw() string {
	return r.raw
}
