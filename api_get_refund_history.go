package appstoreserverapi

import (
	"net/http"
	"sort"
	"strings"
)

// ApiGetRefundHistory 获取退款历史
// Get Refund History
// doc: https://developer.apple.com/documentation/appstoreserverapi/get_refund_history
// desc: true then signedTransactions order by webOrderLineItemId desc
func (c *client) ApiGetRefundHistory(transactionId string, desc bool) (*RefundLookupResponse, error) {
	reqUrl := c.apiGetRefundHistoryUrl + transactionId
	r, err := c.doRequest(http.MethodGet, reqUrl, nil)
	if err != nil {
		return nil, err
	}
	result := &RefundLookupResponse{
		raw: r.String(),
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

type RefundLookupResponse struct {
	raw                string
	SignedTransactions []JWSTransactionDecodedPayload `json:"signedTransactions"`
}

func (r *RefundLookupResponse) Raw() string {
	return r.raw
}
