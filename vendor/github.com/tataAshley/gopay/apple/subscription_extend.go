package apple

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-pay/gopay"
	"net/http"
)

// SubscriptionExtend Extends the renewal date of a customer’s active subscription using the original transaction identifier.
// Doc: https://developer.apple.com/documentation/appstoreserverapi/extend_a_subscription_renewal_date
func (c *Client) SubscriptionExtend(ctx context.Context, transactionId string, bm gopay.BodyMap) (rsp *SubscriptionsExtendRsp, err error) {
	path := fmt.Sprintf(getAllSubscriptionStatuses, transactionId)

	res, bs, err := c.doRequestPut(ctx, path, bm)
	if err != nil {
		return nil, err
	}
	rsp = &SubscriptionsExtendRsp{}
	if err = json.Unmarshal(bs, rsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", errors.New("unmarshal error"), err, string(bs))
	}
	if res.StatusCode == http.StatusOK {
		return rsp, nil
	}
	return rsp, nil
}
