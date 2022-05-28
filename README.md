# appstoreserverapi
App Store Server API(应用商店服务器 API)

## usage | 使用

`go get github.com/lhlyu/appstoreserverapi`

## online | 在线使用

- [link](https://astq-rosy.vercel.app/)

## interface | 接口

```go
type Client interface {

	// ApiGetAllSubscriptionStatuses 获取所有的订阅状态
	// Get All Subscription Statuses
	// doc: https://developer.apple.com/documentation/appstoreserverapi/get_all_subscription_statuses
	ApiGetAllSubscriptionStatuses(transactionId string) (*StatusResponse, error)

	// ApiLookUpOrderId 查找订单 ID
	// Look Up Order ID
	// doc: https://developer.apple.com/documentation/appstoreserverapi/look_up_order_id
	ApiLookUpOrderId(orderId string) (*OrderLookupResponse, error)

	// ApiGetTransactionHistory 获取历史交易记录
	// Get Transaction History
	// doc: https://developer.apple.com/documentation/appstoreserverapi/get_transaction_history
	// desc: true then signedTransactions order by webOrderLineItemId desc
	ApiGetTransactionHistory(transactionId string, desc bool) (*HistoryResponse, error)

	// ApiGetRefundHistory 获取退款历史
	// Get Refund History
	// doc: https://developer.apple.com/documentation/appstoreserverapi/get_refund_history
	// desc: true then signedTransactions order by webOrderLineItemId desc
	ApiGetRefundHistory(transactionId string, desc bool) (*RefundLookupResponse, error)

	// ApiExtendAsubscriptionRenewalDate 延长订阅续订日期
	// Extend a Subscription Renewal Date
	// doc: https://developer.apple.com/documentation/appstoreserverapi/extend_a_subscription_renewal_date
	ApiExtendAsubscriptionRenewalDate(transactionId string, req ExtendRenewalDateRequest) (*ExtendRenewalDateResponse, error)

	// ApiSendConsumptionInformation 发送消费信息
	// Send Consumption Information
	// doc: https://developer.apple.com/documentation/appstoreserverapi/send_consumption_information
	ApiSendConsumptionInformation(transactionId string, req ConsumptionRequest) error
}
```

## example | 例子

```go
func TestClient_GetAllSubscriptionStatuses(t *testing.T) {
	c, err := NewClient(&Config{
		Iss:      ISS,
		Kid:      KID,
		Bid:      BID,
		Pk:       PK,
		Aud:      AUD,
		ExpiryIn: time.Second * 6,
	})
	if err != nil {
		t.Error(err)
		return
	}
	r, err := c.ApiGetAllSubscriptionStatuses("180001239612922")
	if err != nil {
		t.Error(err)
		return
	}
	b, _ := json.Marshal(r)
	fmt.Println(string(b))
}
```

- [more example](./client_test.go)
