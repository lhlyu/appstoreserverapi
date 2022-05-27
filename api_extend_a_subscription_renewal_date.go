package appstoreserverapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// ApiExtendAsubscriptionRenewalDate 延长订阅续订日期
// Extend a Subscription Renewal Date
// doc: https://developer.apple.com/documentation/appstoreserverapi/extend_a_subscription_renewal_date
func (c *client) ApiExtendAsubscriptionRenewalDate(transactionId string, req ExtendRenewalDateRequest) (*ExtendRenewalDateResponse, error) {
	reqUrl := c.api_extend_a_subscription_renewal_date_url + transactionId
	b, _ := json.Marshal(req)
	r, err := c.doRequest(http.MethodPut, reqUrl, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	result := &ExtendRenewalDateResponse{
		raw:                   r.String(),
		EffectiveDate:         r.Get("effectiveDate").Int(),
		OriginalTransactionId: r.Get("originalTransactionId").String(),
		Success:               r.Get("success").Bool(),
		WebOrderLineItemId:    r.Get("webOrderLineItemId").String(),
	}
	return result, err
}

type ExtendRenewalDateRequest struct {
	// 必填。延长订阅续订日期的天数。最大值为 90 天。
	// Required.
	// The number of days to extend the subscription renewal date.
	// The maximum value is 90 days.
	ExtendByDays uint8 `json:"extendByDays"`
	// 必填。订阅日期延长的原因代码。
	// Required.
	// The reason code for the subscription date extension.
	ExtendReasonCode uint8 `json:"extendReasonCode"`
	// 必填。一个字符串，其中包含您提供的用于唯一标识此续订日期扩展请求的值。
	// 字符串的最大长度为 128 个字符。
	// Required.
	// A string that contains a value you provide to uniquely identify this renewal-date-extension request.
	// The maximum length of the string is 128 characters.
	RequestIdentifier string `json:"requestIdentifier"`
}

type ExtendRenewalDateResponse struct {
	raw                   string
	EffectiveDate         int64  `json:"effectiveDate"`
	OriginalTransactionId string `json:"originalTransactionId"`
	Success               bool   `json:"success"`
	WebOrderLineItemId    string `json:"webOrderLineItemId"`
}

func (r *ExtendRenewalDateResponse) Raw() string {
	return r.raw
}
