package gopay

import (
	"context"
	"fmt"
	gopayV2 "github.com/go-pay/gopay"
	"github.com/tataAshley/gopay/apple"
)

type ApplePayService struct {
	client *apple.Client
}

func NewApplePayService() *ApplePayService {
	iss := ""
	bid := ""
	kid := ""
	privateKey := ""

	client, err := apple.NewClient(iss, bid, kid, privateKey, true)
	if err != nil {
		return nil
	}

	return &ApplePayService{
		client: client,
	}
}
func (c *ApplePayService) SubscriptionExtend(ctx context.Context, transactionId, requestIdentifier string) {
	bm := make(gopayV2.BodyMap)
	bm.Set("extendByDays", 30).
		Set("extendReasonCode", 2).
		Set("requestIdentifier", requestIdentifier)
	res, err := c.client.SubscriptionExtend(ctx, transactionId, bm)
	fmt.Println(res)
	fmt.Println(err)
	return
}
