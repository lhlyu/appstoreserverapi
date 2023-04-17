package appstoreserverapi

const (
	productionBaseUrl  = "https://api.storekit.itunes.apple.com"
	developmentBaseUrl = "https://api.storekit-sandbox.itunes.apple.com"
)

const (
	apiGetAllSubscriptionStatusesUri     = "/inApps/v1/subscriptions/"            // + TransactionId
	apiLookupOrderIdUri                  = "/inApps/v1/lookup/"                   // + orderId
	apiGetTransactionHistoryUri          = "/inApps/v1/history/"                  // + TransactionId
	apiGetRefundHistoryUri               = "/inApps/v1/refund/lookup/"            // + TransactionId
	apiExtendASubscriptionRenewalDateUri = "/inApps/v1/subscriptions/extend/"     // + TransactionId
	apiSendConsumptionInformationUri     = "/inApps/v1/transactions/consumption/" // + TransactionId
)
