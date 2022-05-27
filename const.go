package appstoreserverapi

const (
	production_base_url  = "https://api.storekit.itunes.apple.com"
	development_base_url = "https://api.storekit-sandbox.itunes.apple.com"
)

const (
	api_get_all_subscription_statuses_uri      = "/inApps/v1/subscriptions/"            // + TransactionId
	api_lookup_order_id_uri                    = "/inApps/v1/lookup/"                   // + orderId
	api_get_transaction_history_uri            = "/inApps/v1/history/"                  // + TransactionId
	api_get_refund_history_uri                 = "/inApps/v1/refund/lookup/"            // + TransactionId
	api_extend_a_subscription_renewal_date_uri = "/inApps/v1/subscriptions/extend/"     // + TransactionId
	api_send_consumption_information_uri       = "/inApps/v1/transactions/consumption/" // + TransactionId
)
