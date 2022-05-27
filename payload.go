package appstoreserverapi

// JWSTransactionDecodedPayload JWSTransaction解码的有效负载
// doc: https://developer.apple.com/documentation/appstoreserverapi/jwstransactiondecodedpayload
type JWSTransactionDecodedPayload struct {
	AppAccountToken             string `json:"appAccountToken,omitempty"`
	BundleId                    string `json:"bundleId,omitempty"`
	Environment                 string `json:"environment,omitempty"`
	ExpiresDate                 int64  `json:"expiresDate,omitempty"`
	InAppOwnershipType          string `json:"inAppOwnershipType,omitempty"`
	IsUpgraded                  bool   `json:"isUpgraded"`
	OfferIdentifier             string `json:"offerIdentifier"`
	OfferType                   int64  `json:"offerType,omitempty"`
	OriginalPurchaseDate        int64  `json:"originalPurchaseDate,omitempty"`
	OriginalTransactionId       string `json:"originalTransactionId,omitempty"`
	ProductId                   string `json:"productId,omitempty"`
	PurchaseDate                int64  `json:"purchaseDate,omitempty"`
	Quantity                    int64  `json:"quantity,omitempty"`
	RevocationDate              int64  `json:"revocationDate,omitempty"`
	RevocationReason            int64  `json:"revocationReason,omitempty"`
	SignedDate                  int64  `json:"signedDate,omitempty"`
	SubscriptionGroupIdentifier string `json:"subscriptionGroupIdentifier"`
	TransactionId               string `json:"transactionId,omitempty"`
	Type                        string `json:"type,omitempty"`
	WebOrderLineItemId          string `json:"webOrderLineItemId,omitempty"`
}

// JWSRenewalInfoDecodedPayload JWSRenewal信息解码负载
// doc: https://developer.apple.com/documentation/appstoreserverapi/jwsrenewalinfodecodedpayload
type JWSRenewalInfoDecodedPayload struct {
	AutoRenewProductId     string `json:"autoRenewProductId,omitempty"`
	AutoRenewStatus        int64  `json:"autoRenewStatus,omitempty"`
	Environment            string `json:"environment,omitempty"`
	ExpirationIntent       int64  `json:"expirationIntent,omitempty"`
	GracePeriodExpiresDate int64  `json:"gracePeriodExpiresDate,omitempty"`
	IsInBillingRetryPeriod bool   `json:"isInBillingRetryPeriod,omitempty"`
	OfferIdentifier        string `json:"offerIdentifier,omitempty"`
	OfferType              int64  `json:"offerType,omitempty"`
	OriginalTransactionId  string `json:"originalTransactionId,omitempty"`
	PriceIncreaseStatus    int64  `json:"priceIncreaseStatus,omitempty"`
	ProductId              string `json:"productId,omitempty"`
	SignedDate             int64  `json:"signedDate,omitempty"`
}
