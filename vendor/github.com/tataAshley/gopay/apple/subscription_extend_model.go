package apple

type SubscriptionsExtendRsp struct {
	StatusCodeErr
	AppAppleId            int    `json:"appAppleId"`
	EffectiveDate         int64  `json:"effectiveDate"`
	OriginalTransactionId string `json:"originalTransactionId"`
	Success               bool   `json:"success"`
	WebOrderLineItemId    string `json:"webOrderLineItemId"`
}
